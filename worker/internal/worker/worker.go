package worker

import (
	"context"
	"github.com/thyagofr/voting-app/worker/internal/model"
	"github.com/thyagofr/voting-app/worker/internal/pubsub"
	"github.com/thyagofr/voting-app/worker/internal/repository"
	"log/slog"
	"os"
	"time"
)

type Worker struct {
	subscriber pubsub.Subscriber
	repository repository.VoteRepository
	logger     *slog.Logger
}

func NewWorker(subscriber pubsub.Subscriber, repository repository.VoteRepository) *Worker {
	return &Worker{
		subscriber: subscriber,
		repository: repository,
		logger:     slog.New(slog.NewJSONHandler(os.Stdout, nil)).WithGroup("VOTE_WORKER"),
	}
}

func (wk Worker) Do(ctx context.Context) {
	var (
		nWorkers    = 10
		voteChannel = make(chan pubsub.Vote, nWorkers)
	)

	for w := 0; w < nWorkers; w++ {
		select {
		case <-ctx.Done():
			wk.logger.Info("ctx done. finishing worker process...")
			if err := wk.subscriber.Unsubscribe(context.Background()); err != nil {
				wk.logger.Error(err.Error())
			}
			close(voteChannel)
		default:
			wk.subscriber.Subscribe(ctx, voteChannel)
		}
	}

	for vote := range voteChannel {
		go func(voteCtx context.Context, newVote pubsub.Vote) {
			if err := wk.repository.Save(voteCtx, &model.Vote{
				CreationDate: time.Now(),
				Category:     string(newVote.Category),
			}); err != nil {
				wk.logger.Error(err.Error())
			}
		}(ctx, vote)
	}
}
