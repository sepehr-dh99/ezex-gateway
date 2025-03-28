package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezex-io/ezex-gateway/api/graphql"
	"github.com/ezex-io/ezex-gateway/api/graphql/resolver"
	"github.com/ezex-io/ezex-gateway/internal/adapter/grpc/notification"
	"github.com/ezex-io/ezex-gateway/internal/adapter/redis"
	"github.com/ezex-io/ezex-gateway/internal/config"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
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

	redisPort, err := redis.New(cfg.RedisAdapterConfig)
	if err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}

	notificationPort, err := notification.New(cfg.NotificationAdapterConfig)
	if err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}

	authInteractor := auth.NewAuth(cfg.AuthInteractorConfig, logging, notificationPort, redisPort)

	resolve := resolver.NewResolver(authInteractor)

	gql := graphql.New(cfg.GraphqlConfig, resolve, logging, mdl.Recover())

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
