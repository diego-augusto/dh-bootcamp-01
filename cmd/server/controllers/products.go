package controllers

import (
	"arquitetura-go/internal/products"
	"arquitetura-go/pkg/web"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductController {
	return &ProductController{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} request
// @Router /products [get]
func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Router /products [post]
func (c *ProductController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}

		// if req.Name == "" {
		// 	ctx.JSON(
		// 		http.StatusBadRequest,
		// 		web.DecodeError(http.StatusBadRequest, "O nome é obrigatório"),
		// 	)
		// 	return
		// }

		// if req.Type == "" {
		// 	ctx.JSON(
		// 		http.StatusBadRequest,
		// 		web.DecodeError(http.StatusBadRequest, "O tipo é obrigatório"),
		// 	)
		// 	return
		// }

		// if req.Count == 0 {
		// 	ctx.JSON(
		// 		http.StatusBadRequest,
		// 		web.DecodeError(http.StatusBadRequest, "A quantidade é necessária"),
		// 	)
		// 	return
		// }

		// if req.Price == 0 {
		// 	ctx.JSON(
		// 		http.StatusBadRequest,
		// 		web.DecodeError(http.StatusBadRequest, "O preço é obrigatório"),
		// 	)
		// 	return
		// }

		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, p),
		)
	}
}

func (c *ProductController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "O nome do produto é obrigatório"})
			return
		}
		if req.Type == "" {
			ctx.JSON(400, gin.H{"error": "O tipo do produto é obrigatório"})
			return
		}
		if req.Count == 0 {
			ctx.JSON(400, gin.H{"error": "A quantidade é obrigatória"})
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, gin.H{"error": "O preço é obrigatório"})
			return
		}

		p, err := c.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *ProductController) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "O nome do produto é obrigatório"})
			return
		}

		p, err := c.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}

type request struct {
	Name  string  `json:"name" binding:"required"`
	Type  string  `json:"type" binding:"required"`
	Count int     `json:"count" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
