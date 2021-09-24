package google_books

import "books_online_api/business/google_books"

type GoogleBook struct {
	Kind string
	Id  string
	Etag string
	SelfLink string
	VolumeInfo struct {
		Title string
		PublishedDate int
		Description string
		IndustryIdentifiers []struct{
			Type string
			Identifier string
		}
		ReadingModes struct {
			Text  bool
			Image bool
		}
		PageCount int
		PrintType string
		Categories []string
		ImageLinks struct {
			SmallThumbnail string
			Thumbnail string
		}
		Language string
		PreviewLink string
		InfoLink string
		CanonicalVolumeLink string
	}
	SaleInfo struct {
		Country string
		Saleability string
		IsEbook bool
		BuyLink string
	}
	AccessInfo struct {
		Country string
		Viewability string
		Embeddable bool
		PublicDomain bool
		TextToSpeechPermission string
		Epub struct {
			IsAvailable bool
			DownloadLink string
		}
		Pdf struct {
			IsAvailable bool
		}
		WebReaderLink string
		AccessViewStatus string
		QuoteSharingAllowed bool
	}
	SearchInfo struct{
		TextSnippet string
	}
}

func (gb *GoogleBook) ToDomain() google_books.Domain {
	return  google_books.Domain{
		Kind: gb.Kind,
		Id: gb.Kind,
		Etag: gb.Etag,
		SelfLink: gb.SelfLink,
		VolumeInfo: gb.VolumeInfo,
		SaleInfo: gb.SaleInfo,
		AccessInfo: gb.AccessInfo,
		SearchInfo: gb.SearchInfo,
	}
}

func ToListDomain(listGookBooks []GoogleBook) []google_books.Domain {
	result := []google_books.Domain{}

	for _, googleBook := range listGookBooks {
		result = append(result, googleBook.ToDomain())
	}

	return result
}