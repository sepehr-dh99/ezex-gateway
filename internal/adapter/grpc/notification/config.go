package notification

type Config struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

var DefaultConfig = &Config{
	Address: "0.0.0.0",
	Port:    9000,
}

func (c *Config) BasicCheck() error {
	return nil
}
