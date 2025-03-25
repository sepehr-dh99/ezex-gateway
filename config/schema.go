package config

type Config struct {
	Domain        string       `yaml:"domain"`
	Debug         bool         `yaml:"debug"`
	GraphqlServer *GraphServer `yaml:"graphql_server"`
	GRPCClients   []GRPCClient `yaml:"grpc_clients"`
}

type GraphServer struct {
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

type GRPCClient struct {
	Service Service `yaml:"service"`
	Address string  `yaml:"address"`
}
