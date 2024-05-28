package ratelimiter

import (
	"sync"
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

	t.Run("should work correctly with memory storage", func(t *testing.T) {
		storage := NewMemoryTimeStorage()
		limiter := New(2, 1*time.Second, storage)

		var wg sync.WaitGroup
		allowedRequests := 0
		deniedRequests := 0

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if limiter.Allow() {
					allowedRequests++
				} else {
					deniedRequests++
				}
			}()
		}

		wg.Wait()

		assert.Equal(t, allowedRequests, 2, "allowed requests")
		assert.Equal(t, deniedRequests, 8, "denied requests")

		time.Sleep(1 * time.Second)

		allowed := limiter.Allow()
		assert.Equal(t, allowed, true)

	})
}
