package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Kiwiabacaxi/Web-Scraping-go/cmd/dto"
	"github.com/Kiwiabacaxi/Web-Scraping-go/pkg/usecases"
	"github.com/gocolly/colly"
)

func GetQuotes(c *colly.Collector) {
	quotes := usecases.ScrapeQuotes(c)
	saveQuotesToFile(quotes)
}

func saveQuotesToFile(quotes []dto.Quote) {
	jsonQuotes, err := json.MarshalIndent(quotes, "", " ")
	if err != nil {
		fmt.Println("Error marshalling quotes:", err)
		return
	}

	err = os.WriteFile("../../data/quotes.json", jsonQuotes, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
