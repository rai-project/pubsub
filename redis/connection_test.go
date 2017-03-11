package redis

import (
	"testing"

	"github.com/rai-project/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestConnectionStart(t *testing.T) {
	conn, err := New(
		pubsub.Endpoints([]string{
			"localhost:6379",
		}),
	)
	if !assert.NoError(t, err) {
		return
	}
	defer conn.Close()
	assert.NotNil(t, conn)
}
