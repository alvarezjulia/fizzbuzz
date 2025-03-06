package service

import (
	"testing"

	"github.com/alvarezjulia/fizzbuzz/internal/domain"
	"github.com/alvarezjulia/fizzbuzz/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFizzBuzzService_GetStats(t *testing.T) {
	counter := storage.NewRequestCounter()
	service := NewFizzBuzzService(counter)

	// Create test requests
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

	// Process request1 twice and request2 once
	_, _ = service.ProcessFizzBuzz(request1)
	_, _ = service.ProcessFizzBuzz(request1)
	_, _ = service.ProcessFizzBuzz(request2)

	stats := service.GetStats()

	// Verify most frequent request
	assert.Equal(t, request1, stats.Parameters)
	assert.Equal(t, 2, stats.Hits)
}

func TestFizzBuzzService_ProcessFizzBuzz(t *testing.T) {
	tests := []struct {
		name        string
		request     domain.Request
		want        []string
		wantErr     bool
		expectedErr *domain.ErrInvalidRequest
	}{
		{
			name: "valid request",
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 5,
				Limit:         15,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			want: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
			wantErr: false,
		},
		{
			name: "invalid first divisor",
			request: domain.Request{
				FirstDivisor:  -3,
				SecondDivisor: 5,
				Limit:         15,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			wantErr: true,
			expectedErr: &domain.ErrInvalidRequest{
				Field:  "firstDivisor",
				Reason: "must be positive",
			},
		},
		{
			name: "invalid second divisor",
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 0,
				Limit:         15,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			wantErr: true,
			expectedErr: &domain.ErrInvalidRequest{
				Field:  "secondDivisor",
				Reason: "must be positive",
			},
		},
		{
			name: "same divisors",
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 3,
				Limit:         15,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			wantErr: true,
			expectedErr: &domain.ErrInvalidRequest{
				Field:  "divisors",
				Reason: "must be different",
			},
		},
		{
			name: "invalid limit",
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 5,
				Limit:         -1,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			wantErr: true,
			expectedErr: &domain.ErrInvalidRequest{
				Field:  "limit",
				Reason: "must be positive",
			},
		},
		{
			name: "limit exceeds maximum",
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 5,
				Limit:         1000001,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			wantErr: true,
			expectedErr: &domain.ErrInvalidRequest{
				Field:  "limit",
				Reason: "exceeds maximum allowed value",
			},
		},
		{
			name: "empty first word",
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 5,
				Limit:         15,
				FirstWord:     "",
				SecondWord:    "Buzz",
			},
			wantErr: true,
			expectedErr: &domain.ErrInvalidRequest{
				Field:  "firstWord",
				Reason: "cannot be empty",
			},
		},
		{
			name: "empty second word",
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 5,
				Limit:         15,
				FirstWord:     "Fizz",
				SecondWord:    "  ",
			},
			wantErr: true,
			expectedErr: &domain.ErrInvalidRequest{
				Field:  "secondWord",
				Reason: "cannot be empty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := storage.NewRequestCounter()
			service := NewFizzBuzzService(counter)

			got, err := service.ProcessFizzBuzz(tt.request)

			if tt.wantErr {
				require.Error(t, err)
				var errInvalidRequest *domain.ErrInvalidRequest
				if assert.ErrorAs(t, err, &errInvalidRequest) {
					assert.Equal(t, tt.expectedErr.Field, errInvalidRequest.Field)
					assert.Equal(t, tt.expectedErr.Reason, errInvalidRequest.Reason)
				}
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
