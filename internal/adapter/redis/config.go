package redis

import "time"

type Config struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	DB           int           `yaml:"db"`
	Password     string        `yaml:"password"`
	DialTimeout  time.Duration `yaml:"dial_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	PoolSize     int           `yaml:"pool_size"`
	Protocol     int           `yaml:"protocol"`
}

var DefaultConfig = &Config{
	Host:         "localhost",
	Port:         6379,
	DB:           0,
	Password:     "",
	DialTimeout:  5 * time.Second,
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 5 * time.Second,
	PoolSize:     10,
	Protocol:     3,
}

func (c *Config) BasicCheck() error {
	return nil
}
