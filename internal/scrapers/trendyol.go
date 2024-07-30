package scrapers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type TrendyolScraper struct {
	query string
}

func NewTrendyolScraper(query string) *TrendyolScraper {
	return &TrendyolScraper{query: query}
}
func (t *TrendyolScraper) Scrape() []Product {
	var products []Product

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("div.p-card-wrppr", func(e *colly.HTMLElement) {
		product := Product{
			Url:   e.ChildAttr("a", "href"),
			Image: e.ChildAttr("img.p-card-img", "src"),
			Name:  e.ChildAttr("img.p-card-img", "alt"),
			Price: strings.TrimSpace(e.ChildText("div.prc-box-dscntd")),
		}
		products = append(products, product)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraped", r.Request.URL)
	})

	searchUrl := fmt.Sprintf("https://www.trendyol.com/sr?q=%s&qt=%s&st=%s&os=1&sst=PRICE_BY_ASC", t.query, t.query, t.query)
	c.Visit(searchUrl)

	return products
}
