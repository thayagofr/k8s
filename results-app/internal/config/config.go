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

type HTTPServer struct {
	Server
}

func (httpS HTTPServer) DSN() string {
	return fmt.Sprintf("%s:%d", httpS.Host, httpS.Port)
}

type Config struct {
	HTTPServer HTTPServer
	Database   Database
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
		HTTPServer: HTTPServer{
			Server: Server{
				Host: getFromEnvOrDefault("SERVER_HOST", "localhost"),
				Port: getFromEnvOrDefaultInt("SERVER_PORT", 8000),
			},
		},
	}
}
