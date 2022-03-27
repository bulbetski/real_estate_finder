package webscraper

import (
	"fmt"
	"log"
	"real_estate_finder/real_estate_finder/internal/dto"
	"real_estate_finder/real_estate_finder/internal/repository/types"
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
func (ws *webscraper) ParseRentalOffers() error {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	rentalOffers := make([]*dto.RentalOffer, 0)

	c.OnHTML("div.OffersSerpItem__generalInfo", func(e *colly.HTMLElement) {
		offer := &dto.RentalOffer{}
		addr := e.ChildText("div.OffersSerpItem__address")
		offer.FullAddress = addr

		href := e.ChildAttr("a", "href")
		if href != "" && strings.HasPrefix(href, "/offer") {
			offer.Link = rentalOfferPrefix + href
		}

		rentalOffers = append(rentalOffers, offer)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(r.StatusCode)
	})

	for i := 1; i < 13; i++ {
		err := c.Visit(fmt.Sprintf("https://realty.yandex.ru/moskva_i_moskovskaya_oblast/kupit/kvartira/studiya/metro-ramenki/?page=%d", i))
		if err != nil {
			log.Fatalln(err)
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
		return fmt.Errorf("repository.InsertRentalOffers: %v", err)
	}

	return nil
}

func (ws *webscraper) GetRentalOfferse() ([]*dto.RentalOffer, error) {
	rentalOffers, err := ws.repository.GetRentalOffers()
	if err != nil {
		return nil, fmt.Errorf("repository.GetRentalOffers: %v", err)
	}

	return toDtoRentalOffers(rentalOffers), nil
}
