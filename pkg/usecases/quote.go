package usecases

import (
	"sync"

	"github.com/Kiwiabacaxi/Web-Scraping-go/cmd/dto"
	"github.com/Kiwiabacaxi/Web-Scraping-go/pkg/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func ScrapeQuotes(c *colly.Collector) []dto.Quote {
	var wg sync.WaitGroup
	quotes := []model.Quote{}

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		div := e.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		tags := []string{}
		div.Find(".tag").Each(func(i int, s *goquery.Selection) {
			tags = append(tags, s.Text())
		})
		structQuote := model.Quote{
			Text:   quote,
			Author: author,
			Tags:   tags,
		}
		quotes = append(quotes, structQuote)
	})

	c.OnRequest(func(r *colly.Request) {
		wg.Add(1)
	})

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	c.Visit("http://quotes.toscrape.com/")
	wg.Wait()

	quotesDTO := convertToDTO(quotes)
	return quotesDTO
}
