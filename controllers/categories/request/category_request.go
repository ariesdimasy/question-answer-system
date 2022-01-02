package request

type CategoryId struct {
	Id uint `param:"id"`
}

type CreateCategory struct {
	CategoryName string `json:"category_name"`
}

type UpdateCategory struct {
	Id           uint   `param:"id"`
	CategoryName string `json:"category_name"`
}

type DeleteCategory struct {
	Id uint `param:"id"`
}
