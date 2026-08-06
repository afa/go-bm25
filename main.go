package main

import (
	"app/documents"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	// orm

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Text string
}

func main() {
	setup_db()
	e := setup_routes()

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func setup_routes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	docs := e.Group("/documents")
	documents.Routes(docs)

	e.GET("/", rootHandler)
	return e
}

func setup_db() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost dbname=fw_dev sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Document{})
	return db
}

func rootHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "bm"})
}
