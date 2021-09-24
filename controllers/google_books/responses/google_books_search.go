package responses

import "books_online_api/business/google_books"

type GoogleBooksResponse struct {
	Keyword string
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

func FromDomain(gb google_books.Domain) GoogleBooksResponse {
	return GoogleBooksResponse{
		Kind: gb.Kind,
		VolumeInfo: gb.VolumeInfo,
	}
}

func FormListDomain(listGoogleBook []google_books.Domain) []GoogleBooksResponse {
	var result []GoogleBooksResponse
	for _, googleBook := range listGoogleBook {
		result = append(result, FromDomain(googleBook))
	}
	return result
}
