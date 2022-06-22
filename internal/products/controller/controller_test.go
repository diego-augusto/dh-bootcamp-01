package controller

import (
	"arquitetura-go/internal/products/domain"
	"arquitetura-go/internal/products/repository"
	"arquitetura-go/internal/products/service"
	"arquitetura-go/pkg/store"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	_ = os.Setenv("TOKEN", "123456")

	db := store.New(store.FileType, "../../../products.json")

	repo := repository.NewRepository(db)
	service := service.NewService(repo, nil)

	p := NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())

	return r
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

func Test_GetProduct_OK(t *testing.T) {
	// criar um servidor e define suas rotas
	r := createServer()
	// criar uma Request do tipo GET e Response para obter o resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	defer req.Body.Close()

	// diz ao servidor que ele pode atender a solicitação
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []domain.Product
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

func Test_SaveProduct_OK(t *testing.T) {
	// crie o Servidor e defina as Rotas
	r := createServer()
	// crie Request do tipo POST e Response para obter o resultado
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
			"name": "PS4 Pro","type": "Eletrodoméstico","count": 1,"price": 2999
	}`)

	// diga ao servidor que ele pode atender a solicitação
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
