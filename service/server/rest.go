package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	gatePrefix = "/gate/v1"
	rpcPrefix  = "/rpc/v1"
	apiPrefix  = "/api/v2"
)

// Rest -
type Rest struct {
	Router *echo.Echo
}

// RestInit -
func RestInit(e *echo.Echo)  {
	var rest = Rest{
		Router: e,
	}

	rest.Route()
}

// Route -
func (r *Rest) Route() {

	open := r.Router.Group(apiPrefix)

	open.POST("/bung", bung)
}

func bung(c echo.Context) error {
	return c.JSON(http.StatusOK, "normalin normalin!!!")
}
