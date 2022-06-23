package domain

import "context"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	Store(ctx context.Context, id int, name, typee string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(ctx context.Context, id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(ctx context.Context, id int, name string) (Product, error)
	Delete(ctx context.Context, id int) error
}

type ProductService interface {
	GetAll(ctx context.Context) ([]Product, error)
	Store(ctx context.Context, name, typee string, count int, price float64) (Product, error)
	Update(ctx context.Context, id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(ctx context.Context, id int, name string) (Product, error)
	Delete(ctx context.Context, id int) error
}
