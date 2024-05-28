package ratelimiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	t.Run("should allow", func(t *testing.T) {
		storage := NewMemoryTimeStorage()
		limiter := New(2, 1*time.Second, storage)

		assert.Equal(t, limiter.Allow(), true)
	})

	t.Run("should not allow", func(t *testing.T) {
		storage := NewMemoryTimeStorage()
		limiter := New(2, 1*time.Second, storage)

		assert.Equal(t, limiter.Allow(), true)
		assert.Equal(t, limiter.Allow(), true)
		assert.Equal(t, limiter.Allow(), false)
	})

	t.Run("should allow 10_000 request per minute", func(t *testing.T) {
		storage := NewMemoryTimeStorage()
		limiter := New(10_000, 1*time.Minute, storage)

		for i := 0; i < 10_000; i++ {
			assert.Equal(t, limiter.Allow(), true)
		}

		assert.Equal(t, limiter.Allow(), false)
	})
}
