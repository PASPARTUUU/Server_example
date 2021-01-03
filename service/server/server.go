package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// Start -
func Start(e *echo.Echo, addr string) {
	port := fmt.Sprintf(":%v", addr)
	if err := e.Start(port); err != nil {
		log.Printf("[WARN] shutting down the server: %v", err)
	}
}

// Stop -
func Stop(e *echo.Echo, shutdownTimeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("[ERROR] failed to shutdown the http server: %v", err)
		return
	}
	log.Print("[INFO] http server stopped")
}
