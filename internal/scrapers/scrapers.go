package scrapers

type Product struct {
	Source string
	Url    string
	Image  string
	Name   string
	Price  string
}

type Scraper struct {
	products           []Product
	trendyolScraper    *TrendyolScraper
	hespiBuradaScraper *HepsiBuradaScraper
	amazonScraper      *AmazonScraper
}

func New() *Scraper {
	return &Scraper{
		trendyolScraper: NewTrendyolScraper(),
		amazonScraper:   NewAmazonScraper(),
	}
}

func (s *Scraper) Scrape(
	query string, page int,
) []Product {
	trendyolProducts := s.trendyolScraper.Scrape(query, page)
	s.appendProducts(trendyolProducts)

	amazonProducts := s.amazonScraper.Scrape(query, page)
	s.appendProducts(amazonProducts)

	return s.products
}

func (s *Scraper) appendProducts(products []Product) {
	s.products = append(s.products, products...)
}
