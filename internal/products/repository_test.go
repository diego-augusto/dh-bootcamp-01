package products

import (
	"arquitetura-go/pkg/store"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("should return a valid produc list", func(t *testing.T) {

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

		result, _ := repository.GetAll()

		assert.Equal(t, result, input, "should be equal")
	})

	t.Run("should return err when Store returns an error", func(t *testing.T) {

		expectedErr := errors.New("error on connect store / database")

		fileStore := store.MemoryStore{
			ReadMock: func(data interface{}) error {
				return expectedErr
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}

		repository := NewRepository(&fileStore)

		_, err := repository.GetAll()

		assert.Equal(t, err, expectedErr, "should be equal")
	})
}
