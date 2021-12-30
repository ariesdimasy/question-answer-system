package request

type ProductId struct {
	Id uint `json:"id" gorm:"primary_key"`
}

type CategoryId struct {
	CategoryId uint `json:"category_id"`
}

type CreateProduct struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type UpdateProduct struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type DeleteProduct struct {
	Id uint `json:"id" gorm:"primary_key"`
}
