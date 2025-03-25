package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	if err := config.basicCheck(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) basicCheck() error {
	if c.GraphqlServer == nil {
		return errors.New("GraphqlServer is required")
	}

	if c.GraphqlServer.Address == "" {
		return errors.New("graphql_server.address is required")
	}
	if c.GraphqlServer.Port <= 0 || c.GraphqlServer.Port > 65535 {
		return errors.New("graphql_server.port must be between 1 and 65535")
	}
	if c.GraphqlServer.QueryPath == "" {
		return errors.New("graphql_server.query_path is required")
	}

	for i, client := range c.GRPCClients {
		if client.Service == "" {
			return fmt.Errorf("grpc_clients[%d].service.name is required", i)
		}
		if client.Address == "" {
			return fmt.Errorf("grpc_clients[%d].address is required", i)
		}
	}

	return nil
}
