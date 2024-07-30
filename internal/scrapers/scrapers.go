package scrapers

type Scraper struct {
	Trendyol *Trendyol
}

func New(query string) *Scraper {
	return &Scraper{
		Trendyol: NewTrendyolScraper(query),
	}
}

func (s *Scraper) Scrape() {
	s.Trendyol.Scrape()
}
