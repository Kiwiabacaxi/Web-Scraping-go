package usecases

import (
	"github.com/Kiwiabacaxi/Web-Scraping-go/cmd/dto"
	"github.com/Kiwiabacaxi/Web-Scraping-go/pkg/model"
)

func convertToDTO(quotes []model.Quote) []dto.Quote {
	quotesDTO := make([]dto.Quote, len(quotes))
	for i, q := range quotes {
		quotesDTO[i] = dto.Quote{
			Text:   q.Text,
			Author: q.Author,
			Tags:   q.Tags,
		}
	}
	return quotesDTO
}
