package corpora

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Routes(g *echo.Group) {
	g.GET("/", getCorpora)
	g.GET("", getCorpora)
	g.GET("/:id", getCorpus)
	g.POST("/", createCorpus)
	g.POST("", createCorpus)
	g.PUT("/:id", updateCorpus)
	g.DELETE("/:id", deleteCorpus)
}

func getCorpora(c *echo.Context) error {
	return c.JSON(http.StatusOK, []string{"corpus1", "corpus2"})
}
func getCorpus(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"name": "corpus1"})
}
func createCorpus(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
func updateCorpus(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
func deleteCorpus(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
