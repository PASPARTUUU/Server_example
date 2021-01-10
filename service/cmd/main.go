package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PASPARTUUU/Server_example/service/config"
	"github.com/PASPARTUUU/Server_example/service/logger"
	"github.com/PASPARTUUU/Server_example/service/postgres"
	"github.com/PASPARTUUU/Server_example/service/server"
	"github.com/PASPARTUUU/Server_example/tools/errpath"
)

const (
	defaultConfigPath     = "configs/config.toml"
	serverShutdownTimeout = 30 * time.Second
)

func main() {
	fmt.Println("i am alive")

	// Parse flags
	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	// Create logger
	logger := logger.New()

	// Parse configs
	cfg, err := config.Parse(*configPath)
	if err != nil {
		logger.Fatal(errpath.Err(err))
	}

	// connect to postgres
	pgConn, err := postgres.Connect(cfg.Postgres)
	if err != nil {
		logger.Fatal(errpath.Err(err))
	}
	logger.Infoln(errpath.Infof("%+v", pgConn.DB))

	router := server.NewRouter(logger)

	server.RestInit(router)

	go server.Start(router, cfg.ServerPort)
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
