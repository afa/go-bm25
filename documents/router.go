package documents

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Routes(g *echo.Group) {
	g.GET("/:id", getDocument)
	g.GET("/", getDocuments)
	g.GET("", getDocuments)
	g.POST("/", createDocument)
	g.POST("", createDocument)
	g.PUT("/:id", updateDocument)
	g.DELETE("/:id", deleteDocument)
}

func getDocuments(c *echo.Context) error {
	return c.JSON(http.StatusOK, []string{"document1", "document2"})
}

func getDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"document": "document1"})
}
func createDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document created"})
}
func updateDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document updated"})
}
func deleteDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document deleted"})
}
