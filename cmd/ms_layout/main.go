package main

import (
	"flag"
	"os"
	"path/filepath"

	"gitlab.toledo24.ru/web/ms_layout/internal/config"
	"gitlab.toledo24.ru/web/ms_layout/internal/logger"
	"gitlab.toledo24.ru/web/ms_layout/internal/server"
)

func main() {
	parentDir := getProjectPath()

	var configPath string
	flag.StringVar(&configPath, "config", config.PathDefault(parentDir), "path to config file")
	flag.Parse()

	var configs *config.Config = config.MustInit(configPath)

	logPath := getLogPath(parentDir)
	logger := logger.ConfigureLogger(configs.Env, logPath)

	if err := server.Start(configs, logger); err != nil {
		logger.Fatal().Err(err).Msg("server.Start")
	}
}

func getProjectPath() string {
	projectPath := os.Getenv("PROJECT_PATH")
	if projectPath != "" {
		return projectPath
	}

	currentDir, _ := os.Getwd()
	return currentDir
}

func getLogPath(parentDir string) string {
	projectPath := os.Getenv("LOG_PATH")
	if projectPath != "" {
		return projectPath
	}
	return filepath.Join(parentDir, "logs", "local.log")
}
