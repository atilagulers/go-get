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
	Page               int
	products           []Product
	trendyolScraper    *TrendyolScraper
	hespiBuradaScraper *HepsiBuradaScraper
	amazonScraper      *AmazonScraper
}

func New(query string, page int) *Scraper {
	return &Scraper{
		trendyolScraper:    NewTrendyolScraper(query, page),
		hespiBuradaScraper: NewHepsiBuradaScraper(query),
		amazonScraper:      NewAmazonScraper(query, page),
	}
}

func (s *Scraper) Scrape() []Product {
	trendyolProducts := s.trendyolScraper.Scrape()
	s.appendProducts(trendyolProducts)

	amazonProducts := s.amazonScraper.Scrape()
	s.appendProducts(amazonProducts)
	fmt.Println(len(amazonProducts))
	fmt.Println(len(trendyolProducts))
	return s.products
}

func (s *Scraper) appendProducts(products []Product) {
	s.products = append(s.products, products...)
}
