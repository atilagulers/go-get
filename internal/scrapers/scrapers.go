package scrapers

import "fmt"

type Product struct {
	Source string
	Url    string
	Image  string
	Name   string
	Price  string
}

type Scraper struct {
	Query              string
	products           []Product
	trendyolScraper    *TrendyolScraper
	hespiBuradaScraper *HepsiBuradaScraper
}

func New(query string) *Scraper {
	return &Scraper{
		trendyolScraper:    NewTrendyolScraper(query),
		hespiBuradaScraper: NewHepsiBuradaScraper(query),
	}
}

func (s *Scraper) Scrape() []Product {
	// trendyolProducts := s.trendyolScraper.Scrape()
	// s.appendProducts(trendyolProducts)

	hespiBuradaProducts := s.hespiBuradaScraper.Scrape()
	s.appendProducts(hespiBuradaProducts)
	fmt.Println("here", hespiBuradaProducts)
	return s.products
}

func (s *Scraper) appendProducts(products []Product) {
	s.products = append(s.products, products...)
}
