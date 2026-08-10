package main

import (
	"app/corpora"
	"app/documents"
	"app/ranks"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	// orm

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Text  string
	State string
}

type Corpus struct {
	gorm.Model
	Name        string
	ExternalKey string
	State       string
}

func main() {
	setup_db()
	e := setup_routes()

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
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

func setup_routes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	docs := e.Group("/documents")
	documents.Routes(docs)

	corp := e.Group("/corpora")
	corpora.Routes(corp)

	rank := e.Group("/ranks")
	ranks.Routes(rank)

	e.GET("/", rootHandler)
	return e
}

func rootHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "bm"})
}
