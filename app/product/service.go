package product

import (
	"alura-go-base/errors"
	"alura-go-base/types"
	"fmt"
	"net/http"
)

type ProductService struct {
	storage types.IProductRepo
}

func NewProductService(storage types.IProductRepo) *ProductService {
	return &ProductService{storage: storage}
}

func (s *ProductService) CreateProduct(payload types.CreateProductPayload) (*types.Product, *errors.HttpError) {
	p, _ := s.storage.GetProductBySKU(payload.SKU)
	if p != nil {
		return nil, errors.NewHttpError(http.StatusBadRequest, fmt.Sprintf("product already registered for SKU: %v", payload.SKU))
	}

	product := types.Product{
		SKU:         payload.SKU,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       int(payload.Price * 100),
	}

	err := s.storage.CreateProduct(&product)
	if err != nil {
		return nil, errors.New500Error(err)
	}

	return &product, nil
}

func (s *ProductService) GetAllProducts() ([]*types.Product, *errors.HttpError) {
	products, err := s.storage.GetAllProducts()
	if err != nil {
		return nil, errors.New500Error(err)
	}

	return products, nil
}

func (s *ProductService) GetProductBySKU(sku string) (*types.Product, *errors.HttpError) {
	if len(sku) == 0 {
		return nil, errors.NewHttpError(http.StatusBadRequest, "missing 'sku'")
	}

	product, err := s.storage.GetProductBySKU(sku)
	if err != nil {
		return nil, errors.NewHttpError(http.StatusBadRequest, err.Error())

	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id int) *errors.HttpError {
	if id == 0 {
		return errors.NewHttpError(http.StatusBadRequest, "missing path param 'id'")
	}

	err := s.storage.DeleteProduct(id)
	if err != nil {
		return errors.New500Error(err)
	}
	return nil
	// util.WriteJson(w, http.StatusOK, map[string]string{"message": "product deleted with success"})
}
