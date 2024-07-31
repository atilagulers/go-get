package scrapers

import (
	"github.com/gocolly/colly"
)

type HepsiBuradaScraper struct {
	query string
}

func NewHepsiBuradaScraper(query string) *HepsiBuradaScraper {
	return &HepsiBuradaScraper{query: query}
}

func (h *HepsiBuradaScraper) Scrape() []Product {
	var products []Product

	c := colly.NewCollector()

	c.OnHTML("li", func(e *colly.HTMLElement) {

		url := e.ChildAttr("a", "href")
		image := e.ChildAttr("img", "src")
		name := e.ChildText("h3[type='cozy']")
		price := e.ChildText("div[data-test-id='price-current-price']")

		product := Product{
			Url:   url,
			Image: image,
			Name:  name,
			Price: price,
		}
		products = append(products, product)

	})

	return products
}
