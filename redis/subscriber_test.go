package redis

import (
	"bytes"
	"io/ioutil"
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

	err = pub.Publish(channel, bytes.NewBufferString(payload))
	assert.NoError(t, err)

	err = pub.Publish(channel, bytes.NewBufferString(pubsub.EndPayload))
	assert.NoError(t, err)

	err = pub.Stop()
	assert.NoError(t, err)

	msgs := sub.Start()
	for msg := range msgs {
		content, err := ioutil.ReadAll(msg)
		assert.NoError(t, err)
		assert.Equal(t, payload, string(content))
	}
	err = sub.Stop()
	assert.NoError(t, err)
}
