package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvarezjulia/fizzbuzz/internal/domain"
	"github.com/alvarezjulia/fizzbuzz/internal/service"
	"github.com/alvarezjulia/fizzbuzz/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_FizzBuzzHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		request        domain.Request
		expectedStatus int
		expectedBody   domain.Response
		expectError    bool
	}{
		{
			name:   "valid request",
			method: http.MethodPost,
			request: domain.Request{
				FirstDivisor:  3,
				SecondDivisor: 5,
				Limit:         15,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			expectedStatus: http.StatusOK,
			expectedBody: domain.Response{
				Result: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			},
		},
		{
			name:           "invalid method",
			method:         http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
			expectError:    true,
		},
		{
			name:   "invalid request - negative divisor",
			method: http.MethodPost,
			request: domain.Request{
				FirstDivisor:  -3,
				SecondDivisor: 5,
				Limit:         15,
				FirstWord:     "Fizz",
				SecondWord:    "Buzz",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := storage.NewRequestCounter()
			fizzBuzzService := service.NewFizzBuzzService(counter)
			handler := NewHandler(fizzBuzzService)

			body, err := json.Marshal(tt.request)
			require.NoError(t, err)

			req := httptest.NewRequest(tt.method, "/fizzbuzz", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler.FizzBuzzHandler(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectError {
				return
			}

			var response domain.Response
			err = json.NewDecoder(w.Body).Decode(&response)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedBody, response)
		})
	}
}

func TestHandler_StatsHandler(t *testing.T) {
	counter := storage.NewRequestCounter()
	fizzBuzzService := service.NewFizzBuzzService(counter)
	handler := NewHandler(fizzBuzzService)

	// Process some requests first
	request := domain.Request{
		FirstDivisor:  3,
		SecondDivisor: 5,
		Limit:         15,
		FirstWord:     "Fizz",
		SecondWord:    "Buzz",
	}
	_, _ = fizzBuzzService.ProcessFizzBuzz(request)
	_, _ = fizzBuzzService.ProcessFizzBuzz(request)

	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectError    bool
	}{
		{
			name:           "valid request",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid method",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/stats", nil)
			w := httptest.NewRecorder()

			handler.StatsHandler(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectError {
				return
			}

			var stats domain.Stats
			err := json.NewDecoder(w.Body).Decode(&stats)
			require.NoError(t, err)
			assert.Equal(t, request, stats.Parameters)
			assert.Equal(t, 2, stats.Hits)
		})
	}
}
