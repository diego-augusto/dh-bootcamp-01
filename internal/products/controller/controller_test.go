package controller

import (
	"arquitetura-go/internal/products/domain"
	"arquitetura-go/internal/products/domain/mocks"
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	_ = os.Setenv("TOKEN", "123456")

	ginEngine := gin.Default()

	return ginEngine
}

func createRequestTest(
	method string,
	url string,
	body string,
) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAll(t *testing.T) {
	mockProducts := []domain.Product{
		{ID: 1, Name: "Playstation 5", Type: "Eletrônicos", Count: 1, Price: 4399.99},
		{ID: 2, Name: "XBOX Series X", Type: "Eletrônicos", Count: 1, Price: 4500.00},
	}

	mockService := new(mocks.ProductService)
	mockService.On("GetAll", mock.Anything).Return(mockProducts, nil)

	r := createServer()

	NewProduct(r, mockService)

	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAllFail(t *testing.T) {
	mockService := new(mocks.ProductService)
	mockService.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).
		Return([]domain.Product{}, sql.ErrNoRows)

	r := createServer()

	NewProduct(r, mockService)

	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	r.ServeHTTP(rr, req)
}
