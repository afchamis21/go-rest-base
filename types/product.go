package types

import (
	"alura-go-base/errors"
	"net/http"
)

type CreateProductPayload struct {
	SKU         string  `json:"sku" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
}

type Product struct {
	ID          int    `json:"id"`
	SKU         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type IProductRepo interface {
	CreateProduct(product *Product) error
	DeleteProduct(id int) error
	GetProductBySKU(sku string) (*Product, error)
	GetAllProducts() ([]*Product, error)
}

type IProductService interface {
	CreateProduct(product CreateProductPayload) (*Product, *errors.HttpError)
	DeleteProduct(id int) *errors.HttpError
	GetProductBySKU(sku string) (*Product, *errors.HttpError)
	GetAllProducts() ([]*Product, *errors.HttpError)
}

type IProductRouter interface {
	IHandler
	HandleCreateProduct(w http.ResponseWriter, r *http.Request)
	HandleDeleteProduct(w http.ResponseWriter, r *http.Request)
	HandleGetProductBySKU(w http.ResponseWriter, r *http.Request)
	HandleGetAllProducts(w http.ResponseWriter, r *http.Request)
}
