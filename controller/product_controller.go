package controller

import (
	"example.com/auction-api/middleware"
	"example.com/auction-api/model"
	"example.com/auction-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully added product"})
}

func (prd *ProductController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	product, err := prd.ProductService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": product})
}

func (prd *ProductController) Offer(ctx *gin.Context) {
	var productOffer model.ProductOffer

	if err := ctx.ShouldBindJSON(&productOffer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tokenUserId, _ := ctx.Get("id")
	userId, _ := tokenUserId.(float64)

	if err := prd.ProductService.Offer(&productOffer, int(userId)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully offered"})

}

func (prd *ProductController) RegisterProductRoutes(rg *gin.RouterGroup) {
	productRoute := rg.Group("/products")
	productRoute.POST("/", middleware.ValidateToken(prd.RedisService), prd.Add)
	productRoute.GET("/:id", prd.GetById)
	productRoute.PUT("/", middleware.ValidateToken(prd.RedisService), prd.Offer)

}
