package google_books

import "context"

type Domain struct {
		Keyword string
		StartIndex int
		MaxResults int
		Kind string
		Id  string
		Etag string
		SelfLink string
		VolumeInfo struct {
			Title string
			PublishedDate string
			Description string
			IndustryIdentifiers []interface{}
			ReadingModes interface {}
			PageCount int
			PrintType string
			Categories []string
			ImageLinks interface {}
			Language string
			PreviewLink string
			InfoLink string
			CanonicalVolumeLink string
		}
		SaleInfo interface{}
		AccessInfo interface{}
		SearchInfo interface{}
}

type Usecase interface {
	SearchBooks(ctx context.Context, keyword string, startIndex int, maxResult int) ([]Domain, error)
}

type ThirdPartyGoogleBooks interface {
	SearchBooks(ctx context.Context, keyword string, startIndex int, maxResult int) ([]Domain, error)
}