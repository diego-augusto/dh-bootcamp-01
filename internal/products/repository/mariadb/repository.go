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
	products := []domain.Product{}

	rows, err := m.db.Query(sqlGetAll)
	if err != nil {
		return products, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return products, err
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
	product := domain.Product{
		Name:  name,
		Type:  typee,
		Count: count,
		Price: price,
	}

	stmt, err := m.db.Prepare(sqlStore)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	res, err := stmt.Exec(&product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return product, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return product, err
	}

	product.ID = int(lastID)

	return product, nil
}

func (m mariaDBRepository) LastID() (int, error) {
	var maxCount int

	row := m.db.QueryRow(sqlLastID)

	err := row.Scan(&maxCount)
	if err != nil {
		return 0, err
	}

	return maxCount, nil
}

func (m mariaDBRepository) Update(
	id int,
	name, productType string,
	count int,
	price float64,
) (domain.Product, error) {
	product := domain.Product{
		ID:    id,
		Name:  name,
		Type:  productType,
		Count: count,
		Price: price,
	}

	stmt, err := m.db.Prepare(sqlUpdate)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.Exec(
		&product.Name,
		&product.Type,
		&product.Count,
		&product.Price,
		&product.ID,
	)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (m mariaDBRepository) UpdateName(id int, name string) (domain.Product, error) {
	product := domain.Product{ID: id, Name: name}

	stmt, err := m.db.Prepare(sqlUpdateName)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.Exec(&product.Name, &product.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (m mariaDBRepository) Delete(id int) error {
	stmt, err := m.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
