package repository

import (
	"arquitetura-go/internal/products/domain"
	"arquitetura-go/pkg/store"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("should return a valid produc list", func(t *testing.T) {
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

		repository := NewRepository(fileStore)

		result, _ := repository.GetAll(context.Background())

		assert.Equal(t, result, input, "should be equal")
	})

	t.Run("should return err when Store returns an error", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		expectedErr := errors.New("error on connect store / database")

		fileStoreMock := &store.Mock{
			Data: []byte{},
			Err:  expectedErr,
		}

		fileStore.AddMock(fileStoreMock)

		repository := NewRepository(fileStore)

		_, err := repository.GetAll(context.Background())

		// assert.NotNil(t, err, "shouldnt be equal")
		assert.Equal(t, err, expectedErr, "should be equal")
	})
}
