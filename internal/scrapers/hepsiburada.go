package scrapers

import (
	"fmt"

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
	searchUrl := fmt.Sprintf("https://www.hepsiburada.com/ara?q=%s", h.query)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
		colly.AllowedDomains("hepsiburada.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnHTML("li[class^='productListContent']", func(e *colly.HTMLElement) {
		url := e.ChildAttr("a", "href")
		if url != "" {
			fmt.Println("URL:", url)
			image := e.ChildAttr("img", "src")
			name := e.ChildText("h3")
			price := e.ChildText("div[data-test-id='price-current-price']")

			product := Product{
				Url:   url,
				Image: image,
				Name:  name,
				Price: price,
			}
			products = append(products, product)
		}
	})

	err := c.Visit(searchUrl)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return products
}
