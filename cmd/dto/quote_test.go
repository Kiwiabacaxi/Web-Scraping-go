package dto

import (
	"reflect"
	"testing"
)

func TestQuote(t *testing.T) {
	quote := Quote{
		Text:   "Test quote",
		Author: "Test author",
		Tags:   []string{"test1", "test2"},
	}

	if quote.Text != "Test quote" {
		t.Errorf("Expected Text to be 'Test quote', but got '%s'", quote.Text)
	}

	if quote.Author != "Test author" {
		t.Errorf("Expected Author to be 'Test author', but got '%s'", quote.Author)
	}

	expectedTags := []string{"test1", "test2"}
	if !reflect.DeepEqual(quote.Tags, expectedTags) {
		t.Errorf("Expected Tags to be '%v', but got '%v'", expectedTags, quote.Tags)
	}
}
