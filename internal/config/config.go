package config

import (
	"os"

	"github.com/ezex-io/ezex-gateway/api/graphql"
	"github.com/ezex-io/ezex-gateway/internal/adapter/grpc/notification"
	"github.com/ezex-io/ezex-gateway/internal/adapter/redis"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug                     bool                 `yaml:"debug"`
	GraphqlConfig             *graphql.Config      `yaml:"graphql"`
	AuthInteractorConfig      *auth.Config         `yaml:"auth"`
	NotificationAdapterConfig *notification.Config `yaml:"notification_client"`
	RedisAdapterConfig        *redis.Config        `yaml:"redis"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	config.GraphqlConfig = graphql.DefaultConfig
	config.RedisAdapterConfig = redis.DefaultConfig
	config.AuthInteractorConfig = auth.DefaultConfig
	config.NotificationAdapterConfig = notification.DefaultConfig

	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) BasicCheck() error {
	if err := c.GraphqlConfig.BasicCheck(); err != nil {
		return err
	}

	if err := c.NotificationAdapterConfig.BasicCheck(); err != nil {
		return err
	}

	if err := c.AuthInteractorConfig.BasicCheck(); err != nil {
		return err
	}

	return c.RedisAdapterConfig.BasicCheck()
}
