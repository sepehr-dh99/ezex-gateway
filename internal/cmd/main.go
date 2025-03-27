package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezex-io/ezex-gateway/api/graphql"
	"github.com/ezex-io/ezex-gateway/config"
	"github.com/ezex-io/ezex-gateway/internal/auth"
	mdl "github.com/ezex-io/gopkg/middleware/http-mdl"
)

func main() {
	configPath := flag.String("config", "./config.yml", "path to config file")
	flag.Parse()

	logging := slog.Default()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}

	defaultLogLevel := slog.LevelInfo
	if cfg.Debug {
		defaultLogLevel = slog.LevelDebug
	}

	slog.SetLogLoggerLevel(defaultLogLevel)

	a := auth.New()

	gql := graphql.New(cfg.GraphqlServer, logging, a, mdl.Recover())

	gql.Start()
	logging.Info("graphql server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		_ = gql.Stop(context.Background())
		logging.Warn("service interrupted")
	case err := <-gql.Notify():
		logging.Error("graphql server got error", "err", err)
	}
}
