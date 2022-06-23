package repository

import (
	"arquitetura-go/internal/products/domain"
	"context"
	"database/sql"
)

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.ProductRepository {
	return mariaDBRepository{db: db}
}

func (m mariaDBRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	products := []domain.Product{}

	rows, err := m.db.QueryContext(ctx, sqlGetAll)
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
	ctx context.Context,
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

	stmt, err := m.db.PrepareContext(ctx, sqlStore)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	res, err := stmt.ExecContext(ctx, &product.Name, &product.Type, &product.Count, &product.Price)
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
	ctx context.Context,
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

	stmt, err := m.db.PrepareContext(ctx, sqlUpdate)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.ExecContext(
		ctx,
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

func (m mariaDBRepository) UpdateName(ctx context.Context, id int, name string) (domain.Product, error) {
	product := domain.Product{ID: id, Name: name}

	stmt, err := m.db.PrepareContext(ctx, sqlUpdateName)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.ExecContext(ctx, &product.Name, &product.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (m mariaDBRepository) Delete(ctx context.Context, id int) error {
	stmt, err := m.db.PrepareContext(ctx, sqlDelete)
	if err != nil {
		return err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
