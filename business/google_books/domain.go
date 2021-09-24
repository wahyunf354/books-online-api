package google_books

import "context"

type Domain struct {
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

type Usecase interface {
	SearchBooks(ctx context.Context, keyword string) ([]Domain, error)
}

type ThirdPartyGoogleBooks interface {
	SearchBooks(ctx context.Context, keyword string) ([]Domain, error)
}