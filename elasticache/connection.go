package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/k0kubun/pp"
	"github.com/rai-project/aws"
	"github.com/rai-project/config"
	"github.com/rai-project/pubsub"
)

type connection struct {
	options pubsub.Options
}

func New(opts ...pubsub.Option) (*connection, error) {
	options := pubsub.Options{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	var sess *session.Session
	if s, ok := options.Context.Value(sessionKey).(*session.Session); ok && s != nil {
		sess = s.Copy()
	}
	if sess == nil {
		var err error
		sess, err = aws.NewSession()
		if err != nil {
			return nil, err
		}
	}

	if s, ok := options.Context.Value(regionKey).(string); ok && s != "" {
		sess.Config.WithRegion(s)
	}

	if config.IsVerbose || config.IsDebug {
		sess.Config.WithCredentialsChainVerboseErrors(true).WithLogger(log)
	}

	client := elasticache.New(sess)
	out, err := client.DescribeCacheClusters(nil)
	if err != nil {
		return nil, err
	}
	pp.Println(out)

	return &connection{
		options: options,
	}, nil
}
