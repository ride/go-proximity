package radix

import (
	"testing"

	redis "github.com/mediocregopher/radix.v2/pool"
	"github.com/stretchr/testify/assert"
)

func dial(t *testing.T) *redis.Pool {
	client, err := redis.New("tcp", "127.0.0.1:6379", 10)
	assert.Nil(t, err)
	return client
}

const key = "go-proximity:test-set"

func TestZAdd(t *testing.T) {
	c := dial(t)
	wrapper := Wrap(c)
	wrapper.ZAdd(key, 1.0, "test-value")

	results, err := c.Cmd("ZRANGE", key, 0.0, -1.0).List()
	assert.Nil(t, err)
	assert.Equal(t, "test-value", results[0])
	c.Cmd("DEL", key, 0.0, -1.0)
}

func TestZRangeByScore(t *testing.T) {
	c := dial(t)
	wrapper := Wrap(c)
	wrapper.ZAdd(key, 1.0, "test-value")
	results, err := wrapper.ZRangeByScore(key, 0.0, 1.5)
	assert.Nil(t, err)
	assert.Equal(t, "test-value", results[0])
	c.Cmd("DEL", key, 0.0, -1.0)
}
