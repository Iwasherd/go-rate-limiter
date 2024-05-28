package ratelimiter

import (
	"sync"
	"time"
)

type Limiter struct {
	limit    int
	interval time.Duration
	storage  TimeStorage
	mu       sync.Mutex
}

func New(limit int, interval time.Duration, storage TimeStorage) *Limiter {
	return &Limiter{
		limit:    limit,
		interval: interval,
		storage:  storage,
	}
}

func (r *Limiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	// Remove outdated timestamps
	cutoff := now.Add(-r.interval)
	r.storage.RemoveBefore(cutoff)

	if r.storage.Count() < r.limit {
		r.storage.Add(now)
		return true
	}

	return false
}
