package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

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

	_, currentFilePath, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(currentFilePath, "..", "..", "..")

	filePath := filepath.Join(rootDir, "data", "quotes.json")

	err = os.WriteFile(filePath, jsonQuotes, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
