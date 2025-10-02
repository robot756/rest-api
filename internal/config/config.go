package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"connectionString" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IDLETimeout time.Duration `yaml:"idle_timeout" env-default:"55s"`
}

func MustLoad() *Config {
	configPath := "./config/local.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can't read config: %s", err)
	}

	return &cfg
}
