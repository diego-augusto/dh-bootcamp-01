package repository

import (
	"arquitetura-go/internal/products/domain"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mockProducts := []domain.Product{
		{
			ID:    1,
			Name:  "Playstation 5",
			Type:  "Eletrônicos",
			Count: 1,
			Price: 4500,
		},
		{
			ID:    2,
			Name:  "XBOX Series X",
			Type:  "Eletrônicos",
			Count: 1,
			Price: 4500,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "name", "type", "count", "price",
	}).AddRow(
		mockProducts[0].ID,
		mockProducts[0].Name,
		mockProducts[0].Type,
		mockProducts[0].Count,
		mockProducts[0].Price,
	).AddRow(
		mockProducts[1].ID,
		mockProducts[1].Name,
		mockProducts[1].Type,
		mockProducts[1].Count,
		mockProducts[1].Price,
	)

	query := "SELECT \\* FROM products"

	mock.ExpectQuery(query).WillReturnRows(rows)

	productsRepo := NewMariaDBRepository(db)

	result, err := productsRepo.GetAll(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, result[0].Name, "Playstation 5")
	assert.Equal(t, result[1].Name, "XBOX Series X")
}

func TestGetAllFailScan(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "name", "type", "count", "price",
	}).AddRow("", "", "", "", "")

	query := "SELECT \\* FROM products"

	mock.ExpectQuery(query).WillReturnRows(rows)

	productsRepo := NewMariaDBRepository(db)

	_, err = productsRepo.GetAll(context.Background())
	assert.Error(t, err)
}

func TestGetAllFailSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := "SELECT \\* FROM products"

	mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

	productsRepo := NewMariaDBRepository(db)

	_, err = productsRepo.GetAll(context.Background())
	assert.Error(t, err)
}
