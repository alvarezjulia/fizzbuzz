package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/alvarezjulia/fizzbuzz/config"
	"github.com/alvarezjulia/fizzbuzz/internal/handlers"
	"github.com/alvarezjulia/fizzbuzz/internal/service"
	"github.com/alvarezjulia/fizzbuzz/internal/storage"
)

type Server struct {
	port    string
	handler *handlers.Handler
	logger  *slog.Logger
}

func NewServer(config *config.Config) *Server {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	counter := storage.NewRequestCounter()
	fizzBuzzService := service.NewFizzBuzzService(counter)
	handler := handlers.NewHandler(fizzBuzzService)

	return &Server{
		port:    config.Port,
		handler: handler,
		logger:  logger,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/api/fizzbuzz", s.handler.FizzBuzzHandler)
	http.HandleFunc("/api/stats", s.handler.StatsHandler)

	s.logger.Info("Server starting", "port", s.port)
	if err := http.ListenAndServe(":"+s.port, nil); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}
