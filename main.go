package main

import (
	"fmt"
	//"log"
	"net/http"
	"os"

	//"teach/internal/httpclient"

	//cr "teach/internal/config/routers"
	//"teach/internal/config/services"
	"teach/pkg/connector"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

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
	//Implement redis connector

	fmt.Println("All registered routes:")
	data := e.Routes()
	for i := 0; i < len(data); i++ {
		fmt.Printf("Method: %s Path: %s\n", data[i].Method, data[i].Path)
	}
	e.Logger.Fatal(e.Start(":8080"))
}
