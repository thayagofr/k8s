package pubsub

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log/slog"
	"os"
)

type RedisPublisher struct {
	channel string
	client  *redis.Client
	logger  *slog.Logger
}

func NewRedisPublisher(address, password string, channel string) *RedisPublisher {
	return &RedisPublisher{
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

func (pub *RedisPublisher) Publish(ctx context.Context, message Vote) error {
	payload, err := json.Marshal(message)
	if err != nil {
		pub.logger.Error("error marshaling message", slog.Any("message", message))
		return err
	}

	pub.logger.Info("publishing message", slog.Any("message", message))
	if err = pub.client.Publish(ctx, pub.channel, payload).Err(); err != nil {
		pub.logger.Error(
			err.Error(),
			slog.Any("message", message),
		)
		return err
	}
	return nil
}
