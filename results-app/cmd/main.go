package main

import (
	"context"
	"github.com/thyagofr/results-app/internal/config"
	"github.com/thyagofr/results-app/internal/database"
	"github.com/thyagofr/results-app/internal/repository"
	"github.com/thyagofr/results-app/internal/server"
	"net/http"
)

func main() {
	var (
		configurations = config.FromEnv()

		dbConfig     = configurations.Database
		serverConfig = configurations.HTTPServer

		ctx = context.Background()
	)

	db, err := database.OpenConnection(ctx, dbConfig.DSN())
	if err != nil {
		panic(err)
	}

	var voteRepository = repository.NewVotePostgreSQL(db)

	httpServer := server.NewHTTPServer(
		voteRepository,
		"./internal/server/static",
	)

	http.HandleFunc("/", httpServer.Index)
	_ = http.ListenAndServe(serverConfig.DSN(), nil)
}
