package controller

import (
	"example.com/auction-api/middleware"
	"example.com/auction-api/model"
	"example.com/auction-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService  service.UserService
	RedisService service.RedisService
}

func NewUserController(userService service.UserService, redisService service.RedisService) UserController {
	return UserController{
		UserService:  userService,
		RedisService: redisService,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var user model.UserRegister
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully registered"})

}

func (uc *UserController) Login(ctx *gin.Context) {
	var user model.UserLogin
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, err := uc.UserService.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+*token)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged in"})
}

func (uc *UserController) Logout(ctx *gin.Context) {

	userId, _ := ctx.Get("id")
	err := uc.UserService.Logout(fmt.Sprintf("%v", userId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Succesfully logged out"})
}

func (uc *UserController) Validate(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "Validated"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/register", uc.Register)
	userRoute.POST("/login", uc.Login)
	userRoute.POST("/logout", uc.Logout)
	userRoute.GET("/validate", middleware.ValidateToken(uc.RedisService), uc.Validate)

}
