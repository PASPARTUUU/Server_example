package server

import (
	"github.com/labstack/echo/v4"
)

// Option to allow config an echo instance
type Option func(*echo.Echo)

// NewRouter creates a new echo instance, uses provided options and sets the default middleware
func NewRouter() *echo.Echo {
	e := echo.New()

	
	return e
}
