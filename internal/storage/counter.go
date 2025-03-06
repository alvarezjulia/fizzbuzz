package storage

import (
	"fmt"
	"sync"

	"github.com/alvarezjulia/fizzbuzz/internal/domain"
)

type RequestCounter struct {
	sync.Mutex
	counts map[string]int
	params map[string]domain.Request
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{
		counts: make(map[string]int),
		params: make(map[string]domain.Request),
	}
}

func (rc *RequestCounter) UpdateStats(req domain.Request) {
	key := fmt.Sprintf("%d-%d-%d-%s-%s", req.FirstDivisor, req.SecondDivisor, req.Limit, req.FirstWord, req.SecondWord)

	rc.Lock()
	defer rc.Unlock()

	rc.counts[key]++
	rc.params[key] = req
}

func (rc *RequestCounter) GetMostFrequentRequest() domain.Stats {
	rc.Lock()
	defer rc.Unlock()

	var maxHits int
	var mostFreqKey string

	for key, hits := range rc.counts {
		if hits > maxHits {
			maxHits = hits
			mostFreqKey = key
		}
	}

	if maxHits == 0 {
		return domain.Stats{}
	}

	return domain.Stats{
		Parameters: rc.params[mostFreqKey],
		Hits:       maxHits,
	}
}
