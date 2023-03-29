package controller

import (
	"example.com/auction-api/middleware"
	"example.com/auction-api/model"
	"example.com/auction-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		log.Printf(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}
	err := uc.UserService.Register(&user)
	if err != nil {
		log.Printf(err.Error())
		ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully registered"})

}

func (uc *UserController) Login(ctx *gin.Context) {
	//fmt.Printf("%s %s\n", ctx.Request.Body, ctx.Request.Header)
	fmt.Println("-------------")
	fmt.Printf("%s %s\n", ctx.Request.Body, ctx.Request.Header)
	fmt.Println("-------------")
	var user model.UserLogin
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}
	// FE response için user info dön ?
	userResponse, token, err := uc.UserService.Login(&user)
	if err != nil {
		log.Printf(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
		return
	}
	//ctx.SetSameSite(http.SameSiteLaxMode)
	tokenString := "Bearer " + *token
	ctx.SetCookie("access_token", tokenString, 3600*24*30, "", "", false /*because on localhost for now*/, true)
	
	ctx.JSON(http.StatusOK, gin.H{"username": userResponse.FirstName, "id": userResponse.ID})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	var user model.UserLogout
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}
	fmt.Println(user.UserId)
	err := uc.UserService.Logout(user.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
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
