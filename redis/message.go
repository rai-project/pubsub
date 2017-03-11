package redis

import (
	"bytes"
	"io"

	redis "gopkg.in/redis.v5"

	"github.com/rai-project/serializer"
)

type message struct {
	io.Reader
	payload    string
	serializer serializer.Serializer
}

// Done implements github.com/go-kit/kit/pubsub.Message
func (m *message) Done() error {
	return nil
}

func (m *message) Unmarshal(v interface{}) error {
	return m.serializer.Unmarshal([]byte(m.payload), v)
}

// createMessage will convert the amqp.Delivery to a pubsub.Message
func createMessage(serializer serializer.Serializer, d *redis.Message) *message {
	return &message{
		payload:    d.Payload,
		Reader:     bytes.NewBufferString(d.Payload),
		serializer: serializer,
	}
}
