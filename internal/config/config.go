package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ENV        string `yaml:"env" env-required:"prod"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {

	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	pathFile := os.Getenv("pathOfYaml")
	if pathFile == "" {
		log.Fatal("name of yaml file is empty")
	}

	if _, err := os.Stat(pathFile); os.IsNotExist(err) {
		log.Fatalf("File doesn't exist: %s", pathFile)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(pathFile, &cfg); err != nil {
		log.Fatal("cant read config file")
	}

	return &cfg
}
