package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log/slog"
	"os"
)

type RedisSubscriber struct {
	channel    string
	client     *redis.Client
	subscriber *redis.PubSub
	logger     *slog.Logger
}

func NewRedisSubscriber(address, password string, channel string) *RedisSubscriber {
	return &RedisSubscriber{
		client: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
		}),
		logger: slog.New(
			slog.NewJSONHandler(os.Stdout, nil).WithGroup("REDIS_SUBSCRIBER"),
		),
		channel: channel,
	}
}

func (sub *RedisSubscriber) Subscribe(ctx context.Context, outputCh chan Vote) {
	sub.subscriber = sub.client.Subscribe(ctx, sub.channel)
	go func() {
		for {
			msg, err := sub.subscriber.ReceiveMessage(ctx)
			if err != nil {
				sub.logger.Error(
					fmt.Sprintf("error reading message through channel %s", sub.channel),
					slog.String("error", err.Error()),
				)
				continue
			}

			var message Vote
			if err = json.Unmarshal([]byte(msg.Payload), &message); err != nil {
				sub.logger.Error("error unmarshalling message annel", slog.String("error", err.Error()))
				continue
			}
			outputCh <- message
		}
	}()
}

func (sub *RedisSubscriber) Unsubscribe(ctx context.Context) error {
	if sub.subscriber == nil {
		return nil
	}
	return sub.subscriber.Unsubscribe(ctx, sub.channel)
}
