package ratelimiter

import "time"

type TimeStorage interface {
	Add(timestamp time.Time)
	RemoveBefore(cutoff time.Time)
	Count() int
}

type MemoryTimeStorage struct {
	timestamps []time.Time
}

func NewMemoryTimeStorage() *MemoryTimeStorage {
	return &MemoryTimeStorage{}
}

func (s *MemoryTimeStorage) Add(timestamp time.Time) {
	s.timestamps = append(s.timestamps, timestamp)
}

func (s *MemoryTimeStorage) RemoveBefore(cutoff time.Time) {
	var i int
	for i = 0; i < len(s.timestamps); i++ {
		if s.timestamps[i].After(cutoff) {
			break
		}
	}
	s.timestamps = s.timestamps[i:]
}

func (s *MemoryTimeStorage) Count() int {
	return len(s.timestamps)
}
