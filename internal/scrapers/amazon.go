package scrapers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type AmazonScraper struct {
	productPerPage int
}

func NewAmazonScraper() *AmazonScraper {
	return &AmazonScraper{

		productPerPage: 60,
	}
}

func (h *AmazonScraper) Scrape(
	query string, page int,
) []Product {
	var products []Product

	searchUrl := fmt.Sprintf("https://www.amazon.com.tr/s?k=%s&page=%d", query, page)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("div.s-card-container", func(e *colly.HTMLElement) {

		//url := e.ChildAttr("a", "href")
		image := e.ChildAttr("img", "src")
		name := e.ChildText("span.a-text-normal")
		price := strings.Replace(e.ChildText("span.a-price-whole"), ",", "", -1)

		product := Product{
			Source: "Amazon",
			Url:    e.ChildAttr("a.a-link-normal", "href"),
			Image:  image,
			Name:   name,
			Price:  price,
		}

		products = append(products, product)
	})

	err := c.Visit(searchUrl)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return products
}
