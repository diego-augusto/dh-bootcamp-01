package repository

import (
	"arquitetura-go/internal/products/domain"
	"database/sql"
)

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.ProductRepository {
	return mariaDBRepository{db: db}
}

func (m mariaDBRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product

	rows, err := m.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (m mariaDBRepository) Store(
	id int,
	name,
	typee string,
	count int,
	price float64,
) (domain.Product, error) {
	var product domain.Product
	return product, nil
}

func (m mariaDBRepository) LastID() (int, error) {
	return 0, nil
}

func (m mariaDBRepository) Update(
	id int,
	name, productType string,
	count int,
	price float64,
) (domain.Product, error) {
	var product domain.Product
	return product, nil
}

func (m mariaDBRepository) UpdateName(id int, name string) (domain.Product, error) {
	var product domain.Product
	return product, nil
}

func (m mariaDBRepository) Delete(id int) error {
	return nil
}
