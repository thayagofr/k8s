package main

import (
	"github.com/thyagofr/voting-app/internal/config"
	"github.com/thyagofr/voting-app/internal/pubsub"
	"github.com/thyagofr/voting-app/internal/server"
	"net/http"
)

func main() {
	var configurations = config.FromEnv()

	redisPublisher := pubsub.NewRedisPublisher(
		configurations.Redis.DNS(),
		configurations.Redis.Password,
		configurations.Redis.Channel,
	)

	httpServer := server.NewHTTPServer(
		redisPublisher,
		"./internal/server/static",
	)

	http.HandleFunc("/", httpServer.Index)
	_ = http.ListenAndServe(":8080", nil)
}
