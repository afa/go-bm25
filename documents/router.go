package documents
import (
	"net/http"
	"github.com/labstack/echo/v5"
)

func Routes(g *echo.Group) {
	g.GET("/", GetDocuments)
	g.GET("/:id", GetDocument)
	g.POST("/", CreateDocument)
	g.PUT("/:id", UpdateDocument)
	g.DELETE("/:id", DeleteDocument)
	return
}

func GetDocuments(c *echo.Context) error {
	return c.JSON(http.StatusOK, []string{"document1", "document2"})
}

func GetDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"document": "document1"})
}
func CreateDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document created"})
}
func UpdateDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document updated"})
}
func DeleteDocument(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "document deleted"})
}
