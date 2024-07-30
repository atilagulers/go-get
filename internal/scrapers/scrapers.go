package scrapers

type Product struct {
	Url, Image, Name, Price string
}

type Scraper struct {
	Query           string
	products        []Product
	trendyolScraper *TrendyolScraper
}

func New(query string) *Scraper {
	return &Scraper{
		trendyolScraper: NewTrendyolScraper(query),
	}
}

func (s *Scraper) Scrape() []Product {
	trendyolProducts := s.trendyolScraper.Scrape()
	s.appendProducts(trendyolProducts)
	return s.products
}

func (s *Scraper) appendProducts(products []Product) {
	s.products = append(s.products, products...)
}
