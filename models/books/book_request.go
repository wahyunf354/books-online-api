package books

type BookRequest struct {
	Title       string `json:"title"`
	BookTypeId  string `json:"book_type_id"`
	CategoryId  string `json:"category_id"`
	Price       string `json:"price"`
	UserId      string `json:"user_id"`
	Description string `json:"description"`
	PageCount   string `json:"page_count"`
}
