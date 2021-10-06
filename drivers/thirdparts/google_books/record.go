package google_books

import (
	"books_online_api/business/google_books"
)

type GoogleBooks struct {
	Kind string
	TotalItems int
	Items []GoogleBook
}

type GoogleBook struct {
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

func (gb *GoogleBook) ToDomain() google_books.Domain {
	return  google_books.Domain{
		Kind: gb.Kind,
		Id: gb.Id,
		Etag: gb.Etag,
		SelfLink: gb.SelfLink,
		VolumeInfo: gb.VolumeInfo,
		SaleInfo: gb.SaleInfo,
		AccessInfo: gb.AccessInfo,
		SearchInfo: gb.SearchInfo,
	}
}

func ToListDomain(listGookBooks GoogleBooks) []google_books.Domain {
	result := []google_books.Domain{}

	for _, googleBook := range listGookBooks.Items {
		result = append(result, googleBook.ToDomain())
	}

	return result
}