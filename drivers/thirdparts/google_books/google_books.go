package google_books

import (
	"books_online_api/business/google_books"
	"context"
	"fmt"
	"net/http"
)

type ThirtPartGoogleBooks struct {
	LinkAPI string
}

func NewThirtPartGoogleBooks(link string) google_books.ThirdPartyGoogleBooks {
	return &ThirtPartGoogleBooks{
		LinkAPI: link,
	}
}

func (gb *ThirtPartGoogleBooks) SearchBooks(ctx context.Context, keyword string) ([]google_books.Domain, error) {
	var googleBooks []GoogleBook
	result, err := http.Get(gb.LinkAPI+keyword)

	if err != nil {
		return []google_books.Domain{}, err
	}

	defer result.Body.Close()
	body := result.Body

	fmt.Println(body)

	return ToListDomain(googleBooks), nil
}