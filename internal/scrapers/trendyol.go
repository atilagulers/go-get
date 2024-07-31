package scrapers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type TrendyolScraper struct {
	query string
}

func NewTrendyolScraper(query string) *TrendyolScraper {
	return &TrendyolScraper{query: query}
}

func (t *TrendyolScraper) Scrape() []Product {
	var products []Product
	searchUrl := fmt.Sprintf("https://www.trendyol.com/sr?q=%s&qt=%s&st=%s&os=1", t.query, t.query, t.query)

	// Initialize chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Navigate and scrape
	err := chromedp.Run(ctx,
		chromedp.Navigate(searchUrl),
		// Wait for the actual images to load
		chromedp.WaitVisible(`div.p-card-wrppr`, chromedp.ByQuery),
		// Extract product details
		chromedp.Tasks{
			chromedp.ActionFunc(func(ctx context.Context) error {
				var details []map[string]string
				err := chromedp.EvaluateAsDevTools(`[...document.querySelectorAll('div.p-card-wrppr')].map(e => ({
                    url: e.querySelector('a').getAttribute('href'),
                    image: e.querySelector('img.p-card-img').getAttribute('src'),
                    name: e.querySelector('img.p-card-img').getAttribute('alt'),
                    price: e.querySelector('div.prc-box-dscntd').innerText.trim(),
                }))`, &details).Do(ctx)
				if err != nil {
					return err
				}

				for _, d := range details {
					products = append(products, Product{
						Source: "Trendyol",
						Url:    d["url"],
						Image:  d["image"],
						Name:   d["name"],
						Price:  d["price"],
					})
				}
				return nil
			}),
		},
	)
	if err != nil {
		log.Fatalf("Failed to scrape Trendyol: %v", err)
	}

	return products
}
