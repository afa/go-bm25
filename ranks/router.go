package ranks

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Routes(g *echo.Group) {
	g.POST("/simular", rankSimular)
}

func rankSimular(c *echo.Context) error {
	return c.JSON(http.StatusOK, []string{"1", "2"})
}
