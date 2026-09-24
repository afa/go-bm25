package corpora

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	db := c.Get("db").(*gorm.DB)
	list, err := List(db)
	if err != nil {
		panic("list err")
	}
	return c.JSON(http.StatusOK, list)
}

func getCorpusHandler(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	item, err := ByKey(db, id)
	if err != nil {
		panic("item get err")
	}
	return c.JSON(http.StatusOK, item)
}

func createCorpusHandler(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]any{"error": fmt.Sprintf("%w", &err)})
	}
	//json_map has the JSON Payload decoded into a map
	item, err := Create(db, json_map)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]any{"error": fmt.Sprintf("%w", &err)})
	}
	return c.JSON(http.StatusOK, item)
}
func updateCorpusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
func deleteCorpusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
