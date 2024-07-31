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
	amazonScraper      *AmazonScraper
}

func New(query string) *Scraper {
	return &Scraper{
		trendyolScraper:    NewTrendyolScraper(query),
		hespiBuradaScraper: NewHepsiBuradaScraper(query),
		amazonScraper:      NewAmazonScraper(query),
	}
}

func (s *Scraper) Scrape() []Product {
	// trendyolProducts := s.trendyolScraper.Scrape()
	// s.appendProducts(trendyolProducts)

	//hespiBuradaProducts := s.hespiBuradaScraper.Scrape()
	//s.appendProducts(hespiBuradaProducts)

	amazonProducts := s.amazonScraper.Scrape()
	s.appendProducts(amazonProducts)
	fmt.Println(amazonProducts)
	return s.products
}

func (s *Scraper) appendProducts(products []Product) {
	s.products = append(s.products, products...)
}
