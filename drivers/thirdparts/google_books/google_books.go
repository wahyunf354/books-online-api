package google_books

import (
	"books_online_api/business/google_books"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ThirtPartGoogleBooks struct {
	LinkAPI string
}

func NewGoogleBooksAPIThirtPart() google_books.ThirdPartyGoogleBooks {
	return &ThirtPartGoogleBooks{
		LinkAPI: "https://www.googleapis.com/books/v1/volumes?startIndex=%v&maxResults=%v&orderBy=relevance&filter=free-ebooks&q=%v",
	}
}

func (gb *ThirtPartGoogleBooks) SearchBooks(ctx context.Context, keyword string, startIndex int, maxResult int) ([]google_books.Domain, error) {
	var googleBooks GoogleBooks
	link := fmt.Sprintf(gb.LinkAPI, startIndex, maxResult, keyword)
	fmt.Println(link)
	result, err := http.Get(link)

	if err != nil {
		return []google_books.Domain{}, err
	}

	defer result.Body.Close()
	body := result.Body

	if err := json.NewDecoder(body).Decode(&googleBooks); err != nil {
		fmt.Println(err)
		return []google_books.Domain{}, err
	}

	return ToListDomain(googleBooks), nil
}