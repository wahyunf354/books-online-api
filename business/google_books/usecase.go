package google_books

import (
	"context"
	"errors"
	"time"
)

type GoogleBookUserCase struct {
	GoogleBooks ThirdPartyGoogleBooks
	ContextTimeout time.Duration
}

func NewThirtPartGoogleBook(googleBooks ThirdPartyGoogleBooks, timeout time.Duration) Usecase {
	return &GoogleBookUserCase{
		GoogleBooks: googleBooks,
		ContextTimeout: timeout,
	}
}

func (gb *GoogleBookUserCase) SearchBooks (ctx context.Context, keyword string) ([]Domain, error) {
	if keyword != "" {
		return []Domain{}, errors.New("keyword search empty")
	}

	googleBook, err := gb.GoogleBooks.SearchBooks(ctx, keyword)

	return googleBook, err
}


