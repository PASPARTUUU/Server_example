package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/PASPARTUUU/Server_example/server"
	"github.com/PASPARTUUU/Server_example/service/config"
	"github.com/PASPARTUUU/Server_example/service/server"
)

const (
	defaultConfigPath     = "config/myserver.toml"
	maxRequestsAllowed    = 300
	serverShutdownTimeout = 30 * time.Second
	// brokerShutdownTimeout = 30 * time.Second
)

func main() {
	fmt.Println("i am alive")

	router := server.NewRouter()

	server.RestInit(router)

	go server.Start(router, config.ServerPort)
	defer server.Stop(router, serverShutdownTimeout)

	// Wait for program exit
	<-NotifyAboutExit()
}

// NotifyAboutExit -
func NotifyAboutExit() <-chan os.Signal {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	return exit
}
