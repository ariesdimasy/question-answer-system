package request

type CreateCategory struct {
	CategoryName string `json:"category_name"`
}

type UpdateCategory struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}

type DeleteCategory struct {
	Id uint `json:"id"`
}
