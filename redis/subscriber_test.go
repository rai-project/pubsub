package redis

import (
	"testing"

	"github.com/rai-project/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestSubscribeMessage(t *testing.T) {

	conn, err := New(
		pubsub.Endpoints([]string{
			"localhost:6379",
		}),
	)
	if !assert.NoError(t, err) {
		return
	}
	defer conn.Close()

	channel := "test"
	payload := "test message payload"

	sub, err := NewSubscriber(conn, channel)
	assert.NoError(t, err)

	pub, err := NewPublisher(conn)
	assert.NoError(t, err)

	err = pub.Publish(channel, payload)
	assert.NoError(t, err)

	err = pub.End(channel)
	assert.NoError(t, err)

	err = pub.Stop()
	assert.NoError(t, err)

	msgs := sub.Start()
	for msg := range msgs {
		var data string
		err := msg.Unmarshal(&data)
		assert.NoError(t, err)
		assert.Equal(t, payload, data)
	}
	err = sub.Stop()
	assert.NoError(t, err)
}
