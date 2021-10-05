package config

type Config struct {
	AppName string
	IsDebug bool
	Server ServerConfig
}

func New() (*Config, error){
	config := &Config{}

	const defaultPort = 8080

	config.AppName = "counter-api"
	config.Server = ServerConfig{Port: defaultPort}
	config.IsDebug = true

	return config, nil
}