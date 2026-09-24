package corpora

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	g.GET("/", getCorporaHandler)
	g.GET("", getCorporaHandler)
	g.GET("/:id", getCorpusHandler)
	g.POST("/", createCorpusHandler)
	g.POST("", createCorpusHandler)
	g.PUT("/:id", updateCorpusHandler)
	g.DELETE("/:id", deleteCorpusHandler)
}

func getCorporaHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"corpus1", "corpus2"})
}
func getCorpusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"name": "corpus1"})
}
func createCorpusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
func updateCorpusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
func deleteCorpusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
