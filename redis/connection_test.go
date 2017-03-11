package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionStart(t *testing.T) {
	conn, err := New()
	if !assert.NoError(t, err) {
		return
	}
	defer conn.Close()
	assert.NotNil(t, conn)
}
