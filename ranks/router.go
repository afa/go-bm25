package ranks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	g.POST("/simular", rankSimularHandler)
}

func rankSimularHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"1", "2"})
}
