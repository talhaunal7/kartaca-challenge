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
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	server            *gin.Engine
	userService       service.UserService
	redisService      service.RedisService
	userController    controller.UserController
	productService    service.ProductService
	productController controller.ProductController
	ctx               context.Context
	db                *gorm.DB
	err               error
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()

	//dsn := os.Getenv("DB")
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	//db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		panic("failed to migrate the DB")
	}

	redisService = service.NewRedisService(rdb, ctx)
	userService = service.NewUserService(db, redisService)
	userController = controller.NewUserController(userService, redisService)
	productService = service.NewProductService(db)
	productController = controller.NewProductController(productService, redisService)

	server = gin.Default()
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,access_token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	//server.Use(cors.Default())
}
func main() {

	basepath := server.Group("/v1")
	userController.RegisterUserRoutes(basepath)
	productController.RegisterProductRoutes(basepath)
	log.Fatal(server.Run(":3000"))

}
