package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"teach/internal/httpclient"
	"teach/internal/webhook/router"
	"teach/internal/webhook/service"

	cr "teach/internal/config/routers"
	"teach/internal/config/services"
	"teach/pkg/connector"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	wd, _ := os.Getwd()
	if err := godotenv.Load(wd + "/.env"); err != nil {
		fmt.Println("Error loading .env file.")
	}
	fmt.Println("uri = ", os.Getenv("MONGODB_URI"))

	mongodb := connector.NewMongoDBClient(os.Getenv("MONGODB_URI"), 100)
	if mongodb == nil {
		e.Logger.Fatal("Failed to connect to MongoDB")
	}

	db := mongodb.SelectDB(os.Getenv("DATABASE_NAME"))
	if db == nil {
		e.Logger.Fatal("Failed to select database")
	}

	redis := connector.NewRedisClient(
		"",
		os.Getenv("REDIS_ADDRESS"),
		os.Getenv("REDIS_PASSWORD"),
	)
	if _, err := redis.Ping().Result(); err != nil {
		log.Fatalln("Error connecting to Redis")
	}
	h := httpclient.NewHTTPClient(http.Client{}, db, redis)

	webhookService := service.NewWebhookService(db)
	configService := services.NewConfigService(redis, db, h)

	cr.NewConfigRouters(e, configService)

	router.NewWebhookRouter(e, webhookService)
	e.Logger.Fatal(e.Start(":8080"))
}
