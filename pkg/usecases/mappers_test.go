package usecases

import (
	"reflect"
	"testing"

	"github.com/Kiwiabacaxi/Web-Scraping-go/cmd/dto"
	"github.com/Kiwiabacaxi/Web-Scraping-go/pkg/model"
)

func TestConvertToDTO(t *testing.T) {
	// Crie uma slice de model.Quote para testar
	modelQuotes := []model.Quote{
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

	// Execute a função convertToDTO
	dtoQuotes := convertToDTO(modelQuotes)

	// Verifique se a slice resultante tem o mesmo comprimento que a slice original
	if len(dtoQuotes) != len(modelQuotes) {
		t.Errorf("Expected length %d, got %d", len(modelQuotes), len(dtoQuotes))
	}

	// Verifique se a slice resultante tem os mesmos elementos que a slice original
	for i, dtoQuote := range dtoQuotes {
		if !reflect.DeepEqual(dtoQuote, dto.Quote(modelQuotes[i])) {
			t.Errorf("Expected %+v, got %+v", dto.Quote(modelQuotes[i]), dtoQuote)
		}
	}
}
