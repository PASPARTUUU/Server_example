package server

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	// mymware "github.com/PASPARTUUU/Server_example/service/server/middleware"
	mymware "github.com/PASPARTUUU/Server_example/service/server/middleware"
)

// RouterOption to allow config an echo instance
type RouterOption func(*echo.Echo)

// NewRouter - creates a new echo instance, uses provided options and sets the default middleware
func NewRouter(logger *logrus.Logger) *echo.Echo {
	e := echo.New()

	loggerOption(logger)(e)

	return e
}

func loggerOption(logger *logrus.Logger) func(*echo.Echo) {
	return func(e *echo.Echo) {
		e.Logger = &mymware.Logger{Logger: logger} // replace the original echo.Logger with the logrus one
		// Log the requests
		e.Use(mymware.LoggerWithSkipper(
			func(c echo.Context) bool {
				return strings.Contains(c.Request().RequestURI, "/metrics")
			},
		))
	}
}
