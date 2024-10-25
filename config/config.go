package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	PublicHost string
	Port       int

	DbUser    string
	DbPass    string
	DbName    string
	DbPort    int
	DbHost    string
	DbSslMode string
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port:       getEnvInt("PORT", 8080),
		DbUser:     getEnv("DB_USER", "root"),
		DbPass:     getEnv("DB_PASS", "pass"),
		DbName:     getEnv("DB_NAME", "postgres"),
		DbPort:     getEnvInt("DB_PORT", 5432),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbSslMode:  getEnv("DB_SSL_MODE", "disable"),
	}
}

func getEnv(name, fallback string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}

	return fallback
}

func getEnvInt(name string, fallback int) int {
	if val, ok := os.LookupEnv(name); ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		return i
	}

	return fallback
}

func (cfg *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbSslMode)
}

func (cfg *Config) MigrateDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DbUser,    // postgres
		cfg.DbPass,    // postgres
		cfg.DbHost,    // localhost
		cfg.DbPort,    // 5432
		cfg.DbName,    // example
		cfg.DbSslMode, // disable
	)
}
