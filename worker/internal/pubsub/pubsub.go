package pubsub

import "context"

type Subscriber interface {
	Subscribe(ctx context.Context, outputCh chan Vote)
	Unsubscribe(ctx context.Context) error
}
