package controller

import (
	"example.com/auction-api/model"
	"net/http"
	"strconv"

	"example.com/auction-api/middleware"
	"example.com/auction-api/service"
	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := prd.ProductService.Add(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully added product"})
}

func (prd *ProductController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := prd.ProductService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": product})
}

func (prd *ProductController) GetAll(ctx *gin.Context) {

	products, err := prd.ProductService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, _ := range products {
		products[i].User.Password = ""
	}

	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

func (prd *ProductController) Offer(ctx *gin.Context) {
	var productOffer model.ProductOffer
	if err := ctx.ShouldBindJSON(&productOffer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdAny := middleware.GetUserIdFromContext(ctx)
	userId, _ := userIdAny.(float64)
	if err := prd.ProductService.Offer(&productOffer, int(userId)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully offered"})
}

func (prd *ProductController) RegisterProductRoutes(rg *gin.RouterGroup) {
	productRoute := rg.Group("/products")
	productRoute.POST("/", middleware.ValidateToken(prd.RedisService), prd.Add)
	productRoute.GET("/:id", middleware.ValidateToken(prd.RedisService), prd.GetById)
	productRoute.GET("/", middleware.ValidateToken(prd.RedisService), prd.GetAll)
	productRoute.PUT("/offer", middleware.ValidateToken(prd.RedisService), prd.Offer)
}
