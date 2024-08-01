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
	Scrape(query string, page int) []Product
}

func New() *BaseScraper {
	return &BaseScraper{
		trendyolScraper: NewTrendyolScraper(),
		amazonScraper:   NewAmazonScraper(),
	}
}

func (s *BaseScraper) ScrapeAll(
	query string, page, limit int,
) []Product {

	s.scrape(s.trendyolScraper, query, page, limit)
	s.scrape(s.amazonScraper, query, page, limit)

	offset := page * limit

	paginatedProducts := s.products[offset : offset+limit]
	return paginatedProducts

	//return s.products
}

func (s *BaseScraper) scrape(scraper Scraper, query string, page, limit int) {
	startPage, endPage := s.calculatePage(page, limit, s.trendyolScraper.productPerPage)

	for i := startPage; i <= endPage; i++ {
		trendyolProducts := scraper.Scrape(query, i)
		s.appendProducts(trendyolProducts)
	}

}

func (s *BaseScraper) appendProducts(products []Product) {
	s.products = append(s.products, products...)
}

func (s *BaseScraper) calculatePage(offset, limit, productPerPage int) (int, int) {

	fmt.Println("Offset:", offset)
	fmt.Println("Limit:", limit)

	startPage := offset/productPerPage + 1
	endPage := (offset+limit)/productPerPage + 1

	fmt.Println("Start Page:", startPage)
	fmt.Println("End Page:", endPage)

	return startPage, endPage
}
