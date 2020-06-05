package config

type Config struct {
	Endpoint string
}

func GetConfig() Config {
	return Config{
		Endpoint: "localhost:3000",
	}
}
