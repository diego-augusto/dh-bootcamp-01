package main

import (
	"arquitetura-go/docs"
	productsController "arquitetura-go/internal/products/controller"
	productsRepository "arquitetura-go/internal/products/repository/mariadb"
	productsService "arquitetura-go/internal/products/service"
	"database/sql"

	"arquitetura-go/internal/email"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	r := gin.Default()

	// db := store.New(store.FileType, "../../products.json")

	conn, err := sql.Open("mysql", "string de conexão")
	if err != nil {
		log.Fatal("failed to connect to mariadb")
	}

	//1. repositório
	repo := productsRepository.NewMariaDBRepository(conn)

	//2. serviço (regra de negócio)
	//emailSES := email.NewSES()
	emailSendGrid := email.NewSendgrid()
	service := productsService.NewService(repo, emailSendGrid)

	//3. controller
	productsController.NewProduct(r, service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
