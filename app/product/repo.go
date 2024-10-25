package product

import (
	"alura-go-base/types"
	"database/sql"
	"errors"
)

type ProductRepo struct {
	db *sql.DB
}

// GetAllProducts implements types.IProductRepo.
func (p *ProductRepo) GetAllProducts() ([]*types.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		if p.ID == 0 {
			continue
		}

		products = append(products, p)
	}

	return products, nil
}

// CreateProduct implements types.IProductRepo.
func (p *ProductRepo) CreateProduct(product *types.Product) error {
	err := p.db.QueryRow("INSERT INTO products (SKU, Name, Description, Price) VALUES ($1, $2, $3, $4) RETURNING ID",
		product.SKU, product.Name, product.Description, product.Price,
	).Scan(&product.ID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct implements types.IProductRepo.
func (p *ProductRepo) DeleteProduct(id int) error {
	_, err := p.db.Exec("DELETE FROM products p WHERE p.id = $1",
		id,
	)

	return err
}

// GetProductBySKU implements types.IProductRepo.
func (p *ProductRepo) GetProductBySKU(sku string) (*types.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products p WHERE p.sku = $1", sku)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		if p.ID == 0 {
			return nil, errors.New("product not found")
		} else {
			return p, nil
		}

	}

	return nil, errors.New("product not found")
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	p := new(types.Product)

	err := rows.Scan(
		&p.ID,
		&p.SKU,
		&p.Name,
		&p.Description,
		&p.Price,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func NewProductRepo(DB *sql.DB) *ProductRepo {
	return &ProductRepo{db: DB}
}
