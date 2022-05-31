package main

import (
	"arquitetura-go/cmd/server/controllers"
	"arquitetura-go/internal/email"
	"arquitetura-go/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	//1. repositório
	repo := products.NewRepository()

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
