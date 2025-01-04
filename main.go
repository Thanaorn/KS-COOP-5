package main

import (
	"fmt"
	"os"
	"teach/connector"
	"teach/crud/router"
	"teach/crud/service"

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

	fmt.Println("db = ", db)

	CrudService := service.NewCrudService(db)
	router.NewCrudRouter(e, CrudService)

	e.Logger.Fatal(e.Start(":8080"))
}
