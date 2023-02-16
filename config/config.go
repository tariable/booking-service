package config

type Config struct {
	App
	HTTP
	Logger
	Metrics
	Storage
}

type App struct {
	Name    string
	Version string
}

type HTTP struct {
	Host string
	Port string
}

type Logger struct {
}

type Metrics struct {
}

type Storage struct {
}

func GetConfig() (Config, error) {
	/* Hardcoded for the simplicity
	Options:
	1) load config from .env, yml files
	2) load from centralized storage of config (consul, vault)
	*/

	config := Config{
		App{
			Name:    "bookingService",
			Version: "0.0.1",
		},
		HTTP{
			Host: "localhost",
			Port: "8080",
		},
		Logger{},
		Metrics{},
		Storage{},
	}

	return config, nil
}
