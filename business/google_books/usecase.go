package google_books

import (
	"books_online_api/controllers"
	"context"
	"time"
)

type GoogleBookUserCase struct {
	GoogleBooks ThirdPartyGoogleBooks
	ContextTimeout time.Duration
}

func NewGoogleBookThirtPartUsecase(googleBooks ThirdPartyGoogleBooks, timeout time.Duration) Usecase {
	return &GoogleBookUserCase{
		GoogleBooks: googleBooks,
		ContextTimeout: timeout,
	}
}

func (gb *GoogleBookUserCase) SearchBooks (ctx context.Context, keyword string, startIndex int, maxResult int) ([]Domain, error) {
	if keyword == "" {
		return []Domain{}, controllers.KEYWORD_EMPTY
	}

	if maxResult == 0 {
		maxResult = 10
	}

	googleBook, err := gb.GoogleBooks.SearchBooks(ctx, keyword, startIndex, maxResult)

	return googleBook, err
}


