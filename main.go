package main

import (
	"context"
	"example.com/auction-api/controller"
	"example.com/auction-api/entity"
	"example.com/auction-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	server         *gin.Engine
	userService    service.UserService
	redisService   service.RedisService
	userController controller.UserController
	ctx            context.Context
	db             *gorm.DB
	err            error
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()

	dsn := os.Getenv("DB")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to DB")
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("failed to migrate the DB")
	}

	redisService = service.NewRedisService(rdb, ctx)
	userService = service.NewUserService(db, redisService)
	userController = controller.NewUserController(userService, redisService)
	server = gin.Default()

}
func main() {

	basepath := server.Group("/v1")
	userController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":3000"))

}
