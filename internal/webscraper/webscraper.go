package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type repositoryInterface interface {
}

type webscraper struct {
}

func main() {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	// Cache responses to prevent multiple download of pages
	// even if the collector is restarted
	colly.CacheDir("webscraper/cache")

	c.OnHTML("div.OffersSerpItem__generalInfo", func(e *colly.HTMLElement) {
		addr := e.ChildText("div.OffersSerpItem__address")
		fmt.Println(addr)
		href := e.ChildAttr("a", "href")
		if href != "" && strings.HasPrefix(href, "/offer") {
			fmt.Println("https://realty.yandex.ru" + href)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	err := c.Visit("https://realty.yandex.ru/moskva_i_moskovskaya_oblast/kupit/kvartira/studiya/metro-ramenki/")
	if err != nil {
		log.Fatalln(err)
	}
}
