package main

import (
	"arquitetura-go/cmd/server/controllers"
	"arquitetura-go/docs"
	"arquitetura-go/internal/email"
	"arquitetura-go/internal/products"
	"arquitetura-go/pkg/store"
	"arquitetura-go/pkg/web"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set token environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "Token vazio"),
			)
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "Token inválido"),
			)
			return
		}

		c.Next()
	}
}

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

	db := store.New(store.FileType, "../../products.json")

	//1. repositório
	repo := products.NewRepository(db)

	//2. serviço (regra de negócio)
	//emailSES := email.NewSES()
	emailSendGrid := email.NewSendgrid()
	service := products.NewService(repo, emailSendGrid)

	//3. controller
	p := controllers.NewProduct(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	{
		pr.Use(TokenAuthMiddleware())

		pr.POST("/", p.Store())
		pr.GET("/", p.GetAll())
		pr.PUT("/:id", p.Update())
		pr.PATCH("/:id", p.UpdateName())
		pr.DELETE("/:id", p.Delete())
	}

	r.Run()
}
