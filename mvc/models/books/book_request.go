package books

type BookRequest struct {
	Title       string `json:"title" form:"title"`
	BookTypeId  int    `json:"book_type_id" form:"book_type_id"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	Price       int    `json:"price" form:"price"`
	UserId      int    `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
	PageCount   int    `json:"page_count" form:"page_count"`
	ImageUrls   []string
	BookUrl     string `json:"file_book"`
}
