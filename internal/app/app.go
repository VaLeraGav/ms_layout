package app

import (
	"fmt"
	"path/filepath"

	"gitlab.toledo24.ru/web/ms_layout/internal/config"
	"gitlab.toledo24.ru/web/ms_layout/internal/connect_db"
	"gitlab.toledo24.ru/web/ms_layout/internal/logger"
	"gitlab.toledo24.ru/web/ms_layout/internal/server"
	"gitlab.toledo24.ru/web/ms_layout/internal/store/postgres"
)

func Run(config *config.Config, projectPath string) {
	logPath := filepath.Join(projectPath, "logs", "local.log")

	logger := logger.ConfigureLogger(config.Env, logPath)

	psqlInfo := buildConnectUrl(config)
	conn, err := connect_db.New(psqlInfo, config.Db.Driver)
	if err != nil {
		logger.Fatal().Err(err).Msg("connect_db error")
	}

	store := postgres.New(conn)

	var c *server.Server = server.NewServer(config, logger, store)
	c.StartServer()
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
