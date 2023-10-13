package main

import (
	"context"
	"github.com/thyagofr/voting-app/worker/internal/config"
	"github.com/thyagofr/voting-app/worker/internal/database"
	"github.com/thyagofr/voting-app/worker/internal/pubsub"
	"github.com/thyagofr/voting-app/worker/internal/repository"
	"github.com/thyagofr/voting-app/worker/internal/worker"
)

func main() {
	var (
		configurations = config.FromEnv()

		dbConfig    = configurations.Database
		redisConfig = configurations.Redis

		ctx = context.Background()
	)

	db, err := database.OpenConnection(ctx, dbConfig.DSN())
	if err != nil {
		panic(err)
	}

	var (
		redisSubscriber = pubsub.NewRedisSubscriber(redisConfig.DSN(), redisConfig.Password, redisConfig.Channel)
		voteRepository  = repository.NewVotePostgreSQL(db)
		voteWorker      = worker.NewWorker(redisSubscriber, voteRepository)
	)

	voteWorker.Do(ctx)
}
