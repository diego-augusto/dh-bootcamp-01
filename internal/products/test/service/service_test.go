package service_test

import (
	"arquitetura-go/internal/products"
	"arquitetura-go/internal/products/mocks"
	"arquitetura-go/pkg/store"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	t.Run("deve retornar uma lista de produtos ao chamar repository", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		input := []products.Product{
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

		expect := []products.Product{
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
		repository := products.NewRepository(fileStore)

		service := products.NewService(repository, nil)

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
		repository := products.NewRepository(fileStore)

		service := products.NewService(repository, nil)

		_, err := service.GetAll()

		assert.Equal(t, err, expect, "should be equal")
	})
}

func TestGetAll(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)

	p := products.Product{
		ID:    1,
		Name:  "iPhone 13",
		Type:  "Eletrônico",
		Count: 1,
		Price: 5000,
	}

	pList := make([]products.Product, 0)
	pList = append(pList, p)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAll").Return(pList, nil).Once()

		s := products.NewService(mockRepo, nil)
		list, err := s.GetAll()

		assert.NoError(t, err)

		assert.Equal(t, 5000.0, list[0].Price)

		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetAll").
			Return(nil, errors.New("failed to retrieve products")).
			Once()

		s := products.NewService(mockRepo, nil)
		_, err := s.GetAll()

		assert.NotNil(t, err)

		mockRepo.AssertExpectations(t)
	})
}
