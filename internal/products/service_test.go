package products

import (
	"arquitetura-go/pkg/store"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	t.Run("deve retornar uma lista de produtos ao chamar repository", func(t *testing.T) {

		input := []Product{
			{
				ID:    1,
				Name:  "CellPhone",
				Type:  "Tech",
				Count: 3,
				Price: 250,
			}, {
				ID:    2,
				Name:  "Notebook",
				Type:  "Tech",
				Count: 10,
				Price: 1750.5,
			},
		}

		expect := []Product{
			{
				ID:    1,
				Name:  "Tech - CellPhone",
				Type:  "Tech",
				Count: 3,
				Price: 250,
			}, {
				ID:    2,
				Name:  "Tech - Notebook",
				Type:  "Tech",
				Count: 10,
				Price: 1750.5,
			},
		}

		dataJson, _ := json.Marshal(input)

		fileStore := store.MemoryStore{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(dataJson, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}

		repository := NewRepository(&fileStore)

		service := NewService(repository, nil)

		result, _ := service.GetAll()

		assert.Equal(t, result, expect, "should be equal")
	})

	t.Run("deve retornar um error ao chamar repository", func(t *testing.T) {

		expectErr := errors.New("erro ao receber dados")

		fileStore := store.MemoryStore{
			ReadMock: func(data interface{}) error {
				return expectErr
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}

		repository := NewRepository(&fileStore)

		service := NewService(repository, nil)

		_, err := service.GetAll()

		assert.Equal(t, err, expectErr, "should be equal")
	})
}
