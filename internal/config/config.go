package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	HTTPServer `yaml:"http_server" env:"SERVER_PORT"`
	Db         `yaml:"db" env:"SERVER_PORT"`
}

type HTTPServer struct {
	Address string `yaml:"address" env:"HTTP_SERVER_ADDRESS" env-default:"localhost:8080"`
	Timeout int    `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"2"`
}

type Db struct {
	Option   string `yaml:"option" env:"DB_OPTION" `
	Driver   string `yaml:"driver" env:"DB_DRIVER"`
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	NameDb   string `yaml:"name_db" env:"DB_NAME"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
}

func MustInit(configPath string) *Config {
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	err := godotenv.Load(configPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Env: MustGetEnv("ENV"),
		HTTPServer: HTTPServer{
			Address: MustGetEnv("HTTP_SERVER_ADDRESS"),
			Timeout: MustGetEnvAsInt("HTTP_TIMEOUT"),
		},
		Db: Db{
			Option:   MustGetEnv("DB_OPTION"),
			Driver:   MustGetEnv("DB_DRIVER"),
			Host:     MustGetEnv("DB_HOST"),
			Port:     MustGetEnv("DB_PORT"),
			NameDb:   MustGetEnv("DB_NAME"),
			User:     MustGetEnv("DB_USER"),
			Password: MustGetEnv("DB_PASSWORD"),
		},
	}
}

func PathDefault(workDir string) string {
	return filepath.Join(workDir, ".env")
}

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Отсутствует переменая: %s", key)
	}
	return value
}

func MustGetEnvAsInt(name string) int {
	valueStr := MustGetEnv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return -1
}
