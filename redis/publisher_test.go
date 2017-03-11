package redis

import (
	"bytes"
	"testing"

	"github.com/rai-project/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestPublishMessage(t *testing.T) {

	conn, err := New(
		pubsub.Endpoints([]string{
			"localhost:6379",
		}),
	)
	if !assert.NoError(t, err) {
		return
	}
	defer conn.Close()

	pub, err := NewPublisher(conn)
	assert.NoError(t, err)

	err = pub.Publish("test", bytes.NewBufferString("test message payload"))
	assert.NoError(t, err)

	err = pub.Stop()
	assert.NoError(t, err)
}
