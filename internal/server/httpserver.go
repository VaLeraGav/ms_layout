package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.toledo24.ru/web/ms_layout/internal/config"
	"gitlab.toledo24.ru/web/ms_layout/internal/store"
	handlers "gitlab.toledo24.ru/web/ms_layout/internal/ui/handlers"
	"gitlab.toledo24.ru/web/ms_layout/internal/ui/middleware/logger"
	"gitlab.toledo24.ru/web/ms_layout/internal/ui/middleware/request_id"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type Server struct {
	config *config.Config
	router *chi.Mux
	log    *zerolog.Logger
	store  store.Store
}

func NewServer(config *config.Config, logger *zerolog.Logger, store store.Store) *Server {
	s := &Server{
		config: config,
		log:    logger,
		store:  store,
		router: chi.NewRouter(),
	}

	s.configureRouting()

	return s
}

func (s *Server) StartServer() {
	s.log.Log().Str("address", s.config.HTTPServer.Address).Msg("starting server")

	server := &http.Server{
		Addr:    s.config.HTTPServer.Address,
		Handler: s.router,
	}

	shutdownChan := make(chan bool, 1)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.log.Fatal().Err(err).Msg("HTTP server error")
		}

		s.log.Info().Msg("stopped serving new connections")
		shutdownChan <- true
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		s.log.Fatal().Err(err).Msg("HTTP shutdown error")
	}

	<-shutdownChan
	s.log.Info().Msg("graceful shutdown complete")
}

func (s *Server) configureRouting() {
	s.router.Use(request_id.RequestID)
	s.router.Use(logger.New(s.log))

	s.router.Get("/user/{email}", handlers.GetUser(s.log, s.store.User()))
}
