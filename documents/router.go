package documents

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	g.GET("/:id", getDocumentHandler)
	g.GET("/", getDocumentsHandler)
	g.GET("", getDocumentsHandler)
	g.POST("/", createDocumentHandler)
	g.POST("", createDocumentHandler)
	g.PUT("/:id", updateDocumentHandler)
	g.DELETE("/:id", deleteDocumentHandler)
}

func getDocumentsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"document1", "document2"})
}

func getDocumentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"document": "document1"})
}
func createDocumentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document created"})
}
func updateDocumentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document updated"})
}
func deleteDocumentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document deleted"})
}
