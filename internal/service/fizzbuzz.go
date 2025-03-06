package service

import (
	"strconv"
	"strings"

	"github.com/alvarezjulia/fizzbuzz/internal/domain"
	"github.com/alvarezjulia/fizzbuzz/internal/storage"
)

type FizzBuzzService struct {
	counter *storage.RequestCounter
}

func NewFizzBuzzService(counter *storage.RequestCounter) *FizzBuzzService {
	return &FizzBuzzService{
		counter: counter,
	}
}

func (s *FizzBuzzService) ProcessFizzBuzz(req domain.Request) ([]string, error) {
	if err := s.validateRequest(req); err != nil {
		return nil, err
	}

	s.counter.UpdateStats(req)
	return s.generateFizzBuzz(req), nil
}

func (s *FizzBuzzService) GetStats() domain.Stats {
	return s.counter.GetMostFrequentRequest()
}

func (s *FizzBuzzService) validateRequest(req domain.Request) error {
	switch {
	case req.FirstDivisor <= 0:
		return &domain.ErrInvalidRequest{Field: "firstDivisor", Reason: "must be positive"}
	case req.SecondDivisor <= 0:
		return &domain.ErrInvalidRequest{Field: "secondDivisor", Reason: "must be positive"}
	case req.FirstDivisor == req.SecondDivisor:
		return &domain.ErrInvalidRequest{Field: "divisors", Reason: "must be different"}
	case req.Limit <= 0:
		return &domain.ErrInvalidRequest{Field: "limit", Reason: "must be positive"}
	case req.Limit > 1000000:
		return &domain.ErrInvalidRequest{Field: "limit", Reason: "exceeds maximum allowed value"}
	case strings.TrimSpace(req.FirstWord) == "":
		return &domain.ErrInvalidRequest{Field: "firstWord", Reason: "cannot be empty"}
	case strings.TrimSpace(req.SecondWord) == "":
		return &domain.ErrInvalidRequest{Field: "secondWord", Reason: "cannot be empty"}
	}
	return nil
}

func (s *FizzBuzzService) generateFizzBuzz(req domain.Request) []string {
	result := make([]string, req.Limit)
	for i := 1; i <= req.Limit; i++ {
		switch {
		case i%(req.FirstDivisor*req.SecondDivisor) == 0:
			result[i-1] = req.FirstWord + req.SecondWord
		case i%req.FirstDivisor == 0:
			result[i-1] = req.FirstWord
		case i%req.SecondDivisor == 0:
			result[i-1] = req.SecondWord
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
