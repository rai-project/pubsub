package redis

import (
	"context"
	"time"

	"github.com/rai-project/pubsub"
	"gopkg.in/redis.v5"
)

type connection struct {
	*redis.Client
	options pubsub.Options
}

func New(opts ...pubsub.Option) (*connection, error) {
	options := pubsub.Options{
		Endpoints:  Config.Endpoints,
		Password:   Config.Password,
		Serializer: Config.Serializer,
		Context:    context.Background(),
	}

	if Config.Cert != "" {
		pubsub.TLSCertificate(Config.Cert)(&options)
	}

	for _, o := range opts {
		o(&options)
	}

	os := &redis.Options{
		Addr:      options.Endpoints[0],
		Password:  options.Password,
		TLSConfig: options.TLSConfig,
	}

	if val, ok := options.Context.Value(poolsizeKey).(int); ok {
		os.PoolSize = val
	}

	if val, ok := options.Context.Value(poolTimeoutKey).(time.Duration); ok {
		os.PoolTimeout = val
	}

	if val, ok := options.Context.Value(idleTimeoutKey).(time.Duration); ok {
		os.IdleTimeout = val
	}

	if val, ok := options.Context.Value(idleCheckFrequencyKey).(time.Duration); ok {
		os.IdleCheckFrequency = val
	}

	clnt := redis.NewClient(os)

	return &connection{
		Client:  clnt,
		options: options,
	}, nil
}

func (c *connection) Close() error {
	return c.Client.Close()
}

func (c connection) Options() pubsub.Options {
	return c.options
}
