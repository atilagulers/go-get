package scrapers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Trendyol struct {
	Query    string
	Products []TrendyolProduct
}

func NewTrendyolScraper(query string) *Trendyol {
	return &Trendyol{
		Query: query,
	}
}

type TrendyolProduct struct {
	Url, Image, Name, Price string
}

func (t *Trendyol) Scrape() {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("div.p-card-wrppr", func(e *colly.HTMLElement) {
		product := TrendyolProduct{
			Url:   e.ChildAttr("a", "href"),
			Image: e.ChildAttr("img.p-card-img", "src"),
			Name:  e.ChildAttr("img.p-card-img", "alt"),
			Price: strings.TrimSpace(e.ChildText("div.prc-box-dscntd")),
		}
		t.Products = append(t.Products, product)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraped", r.Request.URL)
	})

	searchUrl := fmt.Sprintf("https://www.trendyol.com/sr?q=%s&qt=%s&st=%s&os=1&sst=PRICE_BY_ASC", t.Query, t.Query, t.Query)
	c.Visit(searchUrl)
}
