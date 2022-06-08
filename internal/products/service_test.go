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
		fileStore := store.New(store.FileType, "")

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

		fileStoreMock := &store.Mock{
			Data: dataJson,
			Err:  nil,
		}

		fileStore.AddMock(fileStoreMock)

		//repositório real
		repository := NewRepository(fileStore)

		service := NewService(repository, nil)

		result, _ := service.GetAll()

		assert.Equal(t, result, expect, "should be equal")
	})

	t.Run("deve retornar um error ao chamar repository", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		expect := errors.New("erro ao receber dados")

		fileStoreMock := &store.Mock{
			Data: []byte{},
			Err:  expect,
		}

		fileStore.AddMock(fileStoreMock)

		//repositório real
		repository := NewRepository(fileStore)

		service := NewService(repository, nil)

		_, err := service.GetAll()

		assert.Equal(t, err, expect, "should be equal")
	})
}
