package service

import (
	"arquitetura-go/internal/products/domain"
	"arquitetura-go/internal/products/domain/mocks"
	repository "arquitetura-go/internal/products/repository/file"
	"arquitetura-go/pkg/store"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceGetAll(t *testing.T) {
	t.Run("deve retornar uma lista de produtos ao chamar repository", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		input := []domain.Product{
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

		dataJson, _ := json.Marshal(input)

		fileStoreMock := &store.Mock{
			Data: dataJson,
			Err:  nil,
		}

		fileStore.AddMock(fileStoreMock)

		//repositório real
		repository := repository.NewRepository(fileStore)

		service := NewService(repository, nil)

		result, _ := service.GetAll(context.Background())

		assert.Equal(t, result[0].Name, input[0].Name, "should be equal")
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
		repository := repository.NewRepository(fileStore)

		service := NewService(repository, nil)

		_, err := service.GetAll(context.Background())

		assert.Equal(t, err, expect, "should be equal")
	})
}

func TestGetAll(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)

	p := domain.Product{
		ID:    1,
		Name:  "iPhone 13",
		Type:  "Eletrônico",
		Count: 1,
		Price: 5000,
	}

	pList := make([]domain.Product, 0)
	pList = append(pList, p)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAll").Return(pList, nil).Once()

		s := NewService(mockRepo, nil)
		list, err := s.GetAll(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, 5000.0, list[0].Price)

		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetAll", mock.Anything).
			Return(nil, errors.New("failed to retrieve products")).
			Once()

		s := NewService(mockRepo, nil)
		_, err := s.GetAll(context.Background())

		assert.NotNil(t, err)

		mockRepo.AssertExpectations(t)
	})
}
