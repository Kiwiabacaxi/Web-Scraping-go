package usecases

import (
	"testing"

	"github.com/Kiwiabacaxi/Web-Scraping-go/cmd/dto"
	"github.com/gocolly/colly"
)

func TestScrapeQuotes(t *testing.T) {
	c := colly.NewCollector()

	quotes := ScrapeQuotes(c)

	if len(quotes) == 0 {
		t.Errorf("Expected non-empty slice, got empty slice")
	}

	for _, quote := range quotes {
		if quote.Text == "" {
			t.Errorf("Expected non-empty quote text, got empty string")
		}
		if quote.Author == "" {
			t.Errorf("Expected non-empty quote author, got empty string")
		}
		if len(quote.Tags) == 0 {
			t.Errorf("Expected non-empty quote tags, got empty slice")
		}
	}
}

func MockScrapeQuotes() []dto.Quote {
	return []dto.Quote{
		{
			Text:   "Test quote 1",
			Author: "Test author 1",
			Tags:   []string{"test1", "test2"},
		},
		{
			Text:   "Test quote 2",
			Author: "Test author 2",
			Tags:   []string{"test3", "test4"},
		},
	}
}
