package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DBUrl string
}

func Load() *Config {
	_ = godotenv.Load()

	dbURL := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") +
		"?sslmode=" + os.Getenv("DB_SSLMODE")

	return &Config{
		Port:  os.Getenv("PORT"),
		DBUrl: dbURL,
	}
}

func Must(cfg *Config) {
	if cfg.Port == "" {
		log.Fatal("PORT missing")
	}
}
