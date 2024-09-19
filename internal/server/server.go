package server

import (
	"fmt"

	"gitlab.toledo24.ru/web/ms_layout/internal/config"
	"gitlab.toledo24.ru/web/ms_layout/internal/connect_db"
	"gitlab.toledo24.ru/web/ms_layout/internal/store/postgres"

	"github.com/rs/zerolog"
)

func Start(config *config.Config, logger *zerolog.Logger) error {
	const op = "app.Start"

	psqlInfo := buildConnectUrl(config)
	conn, err := connect_db.New(psqlInfo, config.Db.Driver)
	if err != nil {
		return err
	}

	store := postgres.New(conn)

	var c *Server = NewServer(config, logger, store)
	if err := c.StartServer(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return fmt.Errorf("%s: %w", op, err)
}

func buildConnectUrl(config *config.Config) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s",
		config.Db.Driver,
		config.Db.User,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.NameDb,
		config.Db.Option,
	)
}
