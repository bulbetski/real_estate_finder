package webscraper

import (
	"fmt"
	"log"
	"real_estate_finder/real_estate_finder/internal/dto"
	"real_estate_finder/real_estate_finder/internal/repository/types"
	"regexp"
	"strconv"
	"strings"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/yandex"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const rentalOfferPrefix = "https://realty.yandex.ru"

type repositoryInterface interface {
	InsertRentalOffers(rentalOffers []*types.RentalOffer) error
	GetRentalOffers() ([]*types.RentalOffer, error)
}

type webscraper struct {
	gc         geo.Geocoder
	repository repositoryInterface
}

func New(token string, repository repositoryInterface) *webscraper {
	g := yandex.Geocoder(token)
	return &webscraper{
		gc:         g,
		repository: repository,
	}
}

// ParseRentalOffers ...
// TODO: прокидывать кол-во страниц, которые надо распарсить
func (ws *webscraper) ParseRentalOffers(propertyTypes []types.PropertyType) error {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	rentalOffers := make([]*dto.RentalOffer, 0)

	c.OnHTML("div.OffersSerpItem__info", func(e *colly.HTMLElement) {
		offer := &dto.RentalOffer{}
		generalInfo := e.DOM.Find("div.OffersSerpItem__generalInfo")
		addr := generalInfo.Find("div.OffersSerpItem__address")
		offer.FullAddress = addr.Text()

		href, _ := generalInfo.Find("a").Attr("href")
		if href != "" && strings.HasPrefix(href, "/offer") {
			offer.Link = rentalOfferPrefix + href
		}

		price := strings.ReplaceAll(e.ChildText("div.OffersSerpItem__dealInfo span.price"), "\u00a0", "")
		price = strings.Join(strings.Split(price, " "), "")
		price = strings.TrimSuffix(price, "₽")
		intPrice, err := strconv.ParseInt(price, 10, 64)
		if err != nil {
			// dont panic
		}
		offer.Price = intPrice

		propertyTypeHeader := generalInfo.Find("h3")
		extracted := extractPropertyType(propertyTypeHeader.Text())
		offer.PropertyType = extracted

		rentalOffers = append(rentalOffers, offer)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(r.StatusCode)
	})

	for i := 1; i < 2; i++ {
		url := buildURL(propertyTypes, i)
		err := c.Visit(url)
		if err != nil {
			return fmt.Errorf("failed to visit %s: %v", url, err)
		}
	}

	for _, offer := range rentalOffers {
		loc, err := ws.gc.Geocode(offer.FullAddress)
		if err != nil {
			// value of -1 indicates that there was an error
			offer.Latitude = -1
			offer.Longitude = -1
			continue
		}
		offer.Latitude = loc.Lat
		offer.Longitude = loc.Lng
	}

	err := ws.repository.InsertRentalOffers(toRepoRentalOffers(rentalOffers))
	if err != nil {
		return fmt.Errorf("failed to insert rental offers: %v", err)
	}

	return nil
}

func extractPropertyType(s string) string {
	split := strings.Split(s, ",")
	propertyType := split[len(split)-1]
	if strings.Contains(propertyType, "студия") {
		return types.PropertyType0
	}

	re := regexp.MustCompile("[0-9]+")
	numberOfRooms := re.FindString(propertyType)
	switch numberOfRooms {
	case types.PropertyType1, types.PropertyType2, types.PropertyType3:
		return numberOfRooms
	}

	return types.PropertyType4plus
}

func buildURL(propertyTypes []types.PropertyType, page int) string {
	const baseURL = "https://realty.yandex.ru/moskva_i_moskovskaya_oblast/kupit/kvartira/"
	if len(propertyTypes) == 0 {
		return fmt.Sprintf(baseURL+"?page=%d", page)
	}

	const roomsTotalArg = "roomsTotal="
	var sb strings.Builder
	sb.WriteString(baseURL)
	for i, pType := range propertyTypes {
		if i == 0 {
			sb.WriteString("?")
		}
		sb.WriteString(roomsTotalArg + string(pType))
		if i < len(propertyTypes)-1 {
			sb.WriteString("&")
		}
	}

	return fmt.Sprintf(sb.String()+"&page=%d", page)
}

func (ws *webscraper) GetRentalOfferse() ([]*dto.RentalOffer, error) {
	rentalOffers, err := ws.repository.GetRentalOffers()
	if err != nil {
		return nil, fmt.Errorf("repository.GetRentalOffers: %v", err)
	}

	return toDtoRentalOffers(rentalOffers), nil
}
