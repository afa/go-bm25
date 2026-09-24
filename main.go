package main

import (
	"app/corpora"
	"app/documents"
	"app/ranks"
	"net/http"

	"app/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// orm
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db models.DB

const DBKey = "db"

func ContextDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

func main() {
	Db.DB = setup_db()
	e := setup_routes(Db.DB)

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
	db.AutoMigrate(&models.Document{})
	db.AutoMigrate(&models.Corpus{})
	return db
}

func setup_routes(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(ContextDB(db))

	docs := e.Group("/documents")
	documents.Routes(docs)

	corp := e.Group("/corpora")
	corpora.Routes(corp)

	rank := e.Group("/ranks")
	ranks.Routes(rank)

	e.GET("/", rootHandler)
	return e
}

func rootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "bm"})
}
