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

	_ "github.com/go-sql-driver/mysql"
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

	dataSource := "root:root@tcp(localhost:3306)/bootcamp?parseTime=true"

	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("failed to connect to mariadb")
	}

	// Products domain implementation
	// Repository
	repo := productsRepository.NewMariaDBRepository(conn)

	// Service
	emailSendGrid := email.NewSendgrid()
	service := productsService.NewService(repo, emailSendGrid)

	// Controller
	productsController.NewProduct(r, service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
