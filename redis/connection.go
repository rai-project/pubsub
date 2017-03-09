package redis

import (
	"github.com/rai-project/pubsub"
	"gopkg.in/redis.v5"
)

type connection struct {
	*redis.Client
	options pubsub.Options
}

func New(opts ...pubsub.Option) (*connection, error) {
	options := pubsub.Options{
		Endpoints: Config.Endpoints,
	}

	for _, o := range opts {
		o(&options)
	}

	os := &redis.Options{
		Addr:      options.Endpoints[0],
		Password:  options.Password,
		TLSConfig: options.TLSConfig,
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
