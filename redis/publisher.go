package redis

import (
	"io"
	"io/ioutil"

	"github.com/rai-project/pubsub"
)

type publisher struct {
	conn *connection
}

func NewPublisher(conn *connection) (pubsub.Publisher, error) {
	return &publisher{
		conn: conn,
	}, nil
}

func (p *publisher) Publish(key string, r io.Reader) error {
	conn := p.conn

	content, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return conn.Publish(key, string(content)).Err()
}

func (p *publisher) Stop() error {
	return nil
}
