package redis

import (
	"context"
	"time"

	"github.com/rai-project/pubsub"
)

const (
	databaseKey           = "github.com/rai-project/pubsub/redis/database"
	poolsizeKey           = "github.com/rai-project/pubsub/redis/poolsize"
	poolTimeoutKey        = "github.com/rai-project/pubsub/redis/poolTimeout"
	idleTimeoutKey        = "github.com/rai-project/pubsub/redis/idleTimeout"
	idleCheckFrequencyKey = "github.com/rai-project/pubsub/redis/idleCheckFrequency"
)

func Database(s string) pubsub.Option {
	return func(o *pubsub.Options) {
		o.Context = context.WithValue(o.Context, databaseKey, s)
	}
}

func Poolsize(ii int) pubsub.Option {
	return func(o *pubsub.Options) {
		o.Context = context.WithValue(o.Context, poolsizeKey, ii)
	}
}

func PoolTimeout(t time.Duration) pubsub.Option {
	return func(o *pubsub.Options) {
		o.Context = context.WithValue(o.Context, poolTimeoutKey, t)
	}
}

func IdleTimeout(t time.Duration) pubsub.Option {
	return func(o *pubsub.Options) {
		o.Context = context.WithValue(o.Context, idleTimeoutKey, t)
	}
}

func IdleCheckFrequency(t time.Duration) pubsub.Option {
	return func(o *pubsub.Options) {
		o.Context = context.WithValue(o.Context, idleCheckFrequencyKey, t)
	}
}
