package entity

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.NewString(),
		Name: name,
	}
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}

func NewProduct(name, description string, price float64, categoryID, imageURL string) *Product {
	return &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
	}
}
