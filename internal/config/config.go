package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HttpServer  HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Adress      string        `yaml:"adress" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {

		log.Fatal("CONFIG PATH don't exist")

	}

	if _, err := os.Stat(configPath); err != nil {

		log.Fatalf("config file %s doesn't exist", configPath)

	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("can't read config %s", configPath)
	}

	return &cfg
}
