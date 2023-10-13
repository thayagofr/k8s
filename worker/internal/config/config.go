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

type Redis struct {
	Server
	Credentials
	Channel string
}

func (redis Redis) DSN() string {
	return fmt.Sprintf("%s:%d", redis.Host, redis.Port)
}

type Database struct {
	Credentials
	Server
	Name string
}

func (database Database) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		database.Host,
		database.Port,
		database.Username,
		database.Password,
		database.Name,
	)
}

type Config struct {
	Redis    Redis
	Database Database
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
		Database: Database{
			Credentials: Credentials{
				Username: getFromEnvOrDefault("DB_USERNAME", "test"),
				Password: getFromEnvOrDefault("DB_PASSWORD", "test"),
			},
			Server: Server{
				Host: getFromEnvOrDefault("DB_HOST", "localhost"),
				Port: getFromEnvOrDefaultInt("DB_PORT", 5432),
			},
			Name: getFromEnvOrDefault("DB_NAME", "votes"),
		},
	}
}
