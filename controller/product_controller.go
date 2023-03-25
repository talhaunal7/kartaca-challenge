package controller

import (
	"example.com/auction-api/middleware"
	"example.com/auction-api/model"
	"example.com/auction-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	ProductService service.ProductService
	RedisService   service.RedisService
}

func NewProductController(productService service.ProductService, redisService service.RedisService) ProductController {
	return ProductController{
		ProductService: productService,
		RedisService:   redisService,
	}
}

func (prd *ProductController) Add(ctx *gin.Context) {
	fmt.Println("asdasdasd")
	var product model.ProductAdd

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := prd.ProductService.Add(&product)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

}

func (prd *ProductController) RegisterProductRoutes(rg *gin.RouterGroup) {
	productRoute := rg.Group("/products")
	productRoute.POST("/", middleware.ValidateToken(prd.RedisService), prd.Add)
}
