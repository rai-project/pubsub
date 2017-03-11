package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rai-project/pubsub"
)

const (
	sessionKey = "github.com/rai-project/pubsub/elasticache/session"
	regionKey  = "github.com/rai-project/pubsub/elasticache/region"
)

func Region(s string) pubsub.Option {
	return func(o *pubsub.Options) {
		o.Context = context.WithValue(o.Context, regionKey, s)
	}
}

func Session(s *session.Session) pubsub.Option {
	return func(o *pubsub.Options) {
		o.Context = context.WithValue(o.Context, sessionKey, s)
	}
}
