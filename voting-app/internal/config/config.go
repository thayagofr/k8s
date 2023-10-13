package config

import (
	"fmt"
	"os"
	"strconv"
)

type Credentials struct {
	Username string
	Password string
}

type Server struct {
	Host string
	Port int
}

func (s Server) DNS() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Redis struct {
	Server
	Credentials
	Channel string
}

type Config struct {
	Redis Redis
}

func getFromEnvOrDefaultInt(key string, dftl int) int {
	value := os.Getenv(key)
	converted, err := strconv.Atoi(value)
	if err != nil {
		return dftl
	}
	return converted
}

func getFromEnvOrDefault(key string, dftl string) string {
	value := os.Getenv(key)
	if value == "" {
		return dftl
	}
	return value
}

func FromEnv() *Config {
	return &Config{
		Redis: Redis{
			Credentials: Credentials{
				Username: getFromEnvOrDefault("REDIS_USERNAME", "test"),
				Password: getFromEnvOrDefault("REDIS_PASSWORD", "test"),
			},
			Server: Server{
				Host: getFromEnvOrDefault("REDIS_HOST", "localhost"),
				Port: getFromEnvOrDefaultInt("REDIS_PORT", 6379),
			},
			Channel: getFromEnvOrDefault("REDIS_VOTE_CHANNEL", "vote"),
		},
	}
}
