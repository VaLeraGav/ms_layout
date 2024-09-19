package server

import (
	"net/http"

	"gitlab.toledo24.ru/web/ms_layout/internal/config"
	"gitlab.toledo24.ru/web/ms_layout/internal/store"
	"gitlab.toledo24.ru/web/ms_layout/internal/ui/middleware/logger"
	"gitlab.toledo24.ru/web/ms_layout/internal/ui/middleware/request_id"
	"gitlab.toledo24.ru/web/ms_layout/internal/ui/web"

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

func (s *Server) StartServer() error {
	s.log.Log().Str("address", s.config.HTTPServer.Address).Msg("starting server")

	return http.ListenAndServe(s.config.HTTPServer.Address, s.router)
}

func (s *Server) configureRouting() {
	s.router.Use(request_id.RequestID)
	s.router.Use(logger.New(s.log))

	s.router.Get("/user/{email}", web.GetUser(s.log, s.store.User()))
}
