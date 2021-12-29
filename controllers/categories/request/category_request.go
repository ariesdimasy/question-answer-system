package request

type CreateCategory struct {
	CategoryName string `json:"category_name"`
}

type UpdateCategory struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
}
