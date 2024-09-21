package main

import (
	"flag"
	"os"

	"gitlab.toledo24.ru/web/ms_layout/internal/app"
	"gitlab.toledo24.ru/web/ms_layout/internal/config"
)

func main() {
	currentDir := getProjectPath()

	var configPath string
	flag.StringVar(&configPath, "config", config.PathDefault(currentDir), "path to config file")
	flag.Parse()

	var configs *config.Config = config.MustInit(configPath)

	app.Run(configs, currentDir)
}

func getProjectPath() string {
	projectPath := os.Getenv("PROJECT_PATH")
	if projectPath != "" {
		return projectPath
	}

	currentDir, _ := os.Getwd()
	return currentDir
}
