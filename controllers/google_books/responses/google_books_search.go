package responses

import "books_online_api/business/google_books"

type GoogleBooksResponse struct {
	Keyword string `json:"-"`
	Kind string `json:"kind"`
	Id  string `json:"id"`
	Etag string `json:"etag"`
	SelfLink string `json:"self_link"`
	VolumeInfo struct {
		Title string `json:"title"`
		PublishedDate string `json:"published_date"`
		Description string `json:"description"`
		IndustryIdentifiers []interface{} `json:"industry_identifiers"`
		ReadingModes interface {} `json:"reading_modes"`
		PageCount int `json:"page_count"`
		PrintType string `json:"print_type"`
		Categories []string `json:"categories"`
		ImageLinks interface {} `json:"image_links"`
		Language string `json:"language"`
		PreviewLink string `json:"preview_link"`
		InfoLink string `json:"info_link"`
		CanonicalVolumeLink string `json:"canonical_volume_link"`
	} `json:"volume_info"`
	SaleInfo interface {} `json:"sale_info"`
	AccessInfo interface {} `json:"access_info"`
	SearchInfo interface{} `json:"search_info"`
}

func FromDomain(gb google_books.Domain) GoogleBooksResponse {
	return GoogleBooksResponse{
		Keyword: gb.Keyword,
		Kind: gb.Kind,
		SelfLink: gb.SelfLink,
		Etag: gb.Etag,
		Id: gb.Id,
		VolumeInfo: struct {
			Title               string        `json:"title"`
			PublishedDate       string        `json:"published_date"`
			Description         string        `json:"description"`
			IndustryIdentifiers []interface{} `json:"industry_identifiers"`
			ReadingModes        interface{}   `json:"reading_modes"`
			PageCount           int           `json:"page_count"`
			PrintType           string        `json:"print_type"`
			Categories          []string      `json:"categories"`
			ImageLinks          interface{}   `json:"image_links"`
			Language            string        `json:"language"`
			PreviewLink         string        `json:"preview_link"`
			InfoLink            string        `json:"info_link"`
			CanonicalVolumeLink string        `json:"canonical_volume_link"`
		}(gb.VolumeInfo),
		SaleInfo: gb.SaleInfo,
		AccessInfo: gb.AccessInfo,
		SearchInfo: gb.SearchInfo,
	}
}

func FormListDomain(listGoogleBook []google_books.Domain) []GoogleBooksResponse {
	var result []GoogleBooksResponse
	for _, googleBook := range listGoogleBook {
		result = append(result, FromDomain(googleBook))
	}
	return result
}
