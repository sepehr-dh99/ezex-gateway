package graphql

type Config struct {
	Address    string `yaml:"address"`
	Port       int    `yaml:"port"`
	Playground bool   `yaml:"playground"`
	QueryPath  string `yaml:"query_path"`
	CORS       Cors   `yaml:"cors"`
}

type Cors struct {
	AllowedOrigins   []string `yaml:"allowed_origins"`
	AllowedMethods   []string `yaml:"allowed_methods"`
	AllowedHeaders   []string `yaml:"allowed_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
}

var DefaultConfig = &Config{
	Address:    "0.0.0.0",
	Port:       8080,
	Playground: true,
	QueryPath:  "",
	CORS: Cors{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	},
}

func (c *Config) BasicCheck() error {
	return nil
}
