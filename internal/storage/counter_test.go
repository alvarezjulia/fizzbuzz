package storage

import (
	"sync"
	"testing"

	"github.com/alvarezjulia/fizzbuzz/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRequestCounter_UpdateStats(t *testing.T) {
	counter := NewRequestCounter()

	request := domain.Request{
		FirstDivisor:  3,
		SecondDivisor: 5,
		Limit:         15,
		FirstWord:     "Fizz",
		SecondWord:    "Buzz",
	}

	// Test concurrent updates
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.UpdateStats(request)
		}()
	}
	wg.Wait()

	stats := counter.GetMostFrequentRequest()
	assert.Equal(t, request, stats.Parameters)
	assert.Equal(t, 100, stats.Hits)
}

func TestRequestCounter_GetMostFrequentRequest(t *testing.T) {
	counter := NewRequestCounter()

	request1 := domain.Request{
		FirstDivisor:  3,
		SecondDivisor: 5,
		Limit:         15,
		FirstWord:     "Fizz",
		SecondWord:    "Buzz",
	}

	request2 := domain.Request{
		FirstDivisor:  2,
		SecondDivisor: 4,
		Limit:         10,
		FirstWord:     "Even",
		SecondWord:    "Four",
	}

	// Update stats with different frequencies
	for i := 0; i < 3; i++ {
		counter.UpdateStats(request1)
	}
	for i := 0; i < 2; i++ {
		counter.UpdateStats(request2)
	}

	stats := counter.GetMostFrequentRequest()
	assert.Equal(t, request1, stats.Parameters)
	assert.Equal(t, 3, stats.Hits)
}

func TestRequestCounter_EmptyStats(t *testing.T) {
	counter := NewRequestCounter()
	stats := counter.GetMostFrequentRequest()
	assert.Equal(t, domain.Stats{}, stats)
}
