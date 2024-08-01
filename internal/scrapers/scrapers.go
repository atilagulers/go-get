package scrapers

import "fmt"

type Product struct {
	Source string
	Url    string
	Image  string
	Name   string
	Price  string
}

type BaseScraper struct {
	products           []Product
	trendyolScraper    *TrendyolScraper
	hespiBuradaScraper *HepsiBuradaScraper
	amazonScraper      *AmazonScraper
}

type Scraper interface {
	Scrape(query string, page int, sort string) []Product
}

func New() *BaseScraper {
	return &BaseScraper{
		trendyolScraper: NewTrendyolScraper(),
		amazonScraper:   NewAmazonScraper(),
	}
}

func (s *BaseScraper) ScrapeAll(
	query string, offset, limit int, sort string,
) []Product {

	//s.scrape(s.trendyolScraper, query, offset, limit, sort)
	s.scrape(s.amazonScraper, query, offset, limit, sort)

	paginatedProducts := s.products[offset : offset+limit]

	for _, product := range paginatedProducts {
		fmt.Printf("Price: %s\n", product.Price)
	}

	return paginatedProducts

}

func (s *BaseScraper) scrape(scraper Scraper, query string, offset, limit int, sort string) {
	startPage, endPage := s.calculatePage(offset, limit, s.trendyolScraper.productPerPage)

	for i := startPage; i <= endPage; i++ {
		scraper := scraper.Scrape(query, i, sort)
		s.appendProducts(scraper)
	}
}

func (s *BaseScraper) appendProducts(products []Product) {
	s.products = append(s.products, products...)
}

func (s *BaseScraper) calculatePage(offset, limit, productPerPage int) (int, int) {

	startPage := offset/productPerPage + 1
	endPage := (offset+limit)/productPerPage + 1

	return startPage, endPage
}
