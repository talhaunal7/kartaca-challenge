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
		log.Fatal(err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ctx = context.Background()

	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		log.Fatal(err.Error())
	}

	redisService = service.NewRedisService(rdb, ctx)
	userService = service.NewUserService(db, redisService)
	userController = controller.NewUserController(userService, redisService)
	productService = service.NewProductService(db)
	productController = controller.NewProductController(productService, redisService)

	server = gin.Default()
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,access_token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
}

func main() {
	initData()

	basepath := server.Group("/v1")
	userController.RegisterUserRoutes(basepath)
	productController.RegisterProductRoutes(basepath)
	log.Fatal(server.Run(":8080"))
}

func initData() {
	var count int64
	db.Model(&entity.Product{}).Count(&count)

	if count == 0 {
		products := []entity.Product{
			{
				Name:       "Iznik Design Ceramic Vase",
				OfferPrice: 1,
				ImgUrl:     "https://imgur.com/KP1co0W.jpg",
			},
			{
				Name:       "Handmade Artifact Necklace",
				OfferPrice: 1,
				ImgUrl:     "https://imgur.com/pyvM57E.jpg",
			},
			{
				Name:       "Persian Carpet - Gordes Carpet",
				OfferPrice: 1,
				ImgUrl:     "https://imgur.com/0UK3CdX.jpg",
			},
		}

		if err = db.Create(&products).Error; err != nil {
			log.Fatal(err.Error())
		}
	}
}
