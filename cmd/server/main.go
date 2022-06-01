package main

import (
	"arquitetura-go/cmd/server/controllers"
	"arquitetura-go/internal/email"
	"arquitetura-go/internal/products"
	"arquitetura-go/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
