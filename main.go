package main

import (
	// "fmt"
	"app/documents"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	docs := e.Group("documents")
	documents.Routes(docs)
	
	e.GET("/", rootHandler)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func rootHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "bm"})
}
