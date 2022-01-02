package request

type ProductId struct {
	Id uint `param:"category_id"`
}

type CategoryId struct {
	CategoryId uint `param:"category_id"`
}

type CreateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type UpdateProduct struct {
	Id          uint   `param:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type DeleteProduct struct {
	Id uint `param:"id"`
}
