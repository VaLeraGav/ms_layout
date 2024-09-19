package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gitlab.toledo24.ru/web/ms_layout/internal/config"
	"gitlab.toledo24.ru/web/ms_layout/internal/connect_db"
	"gitlab.toledo24.ru/web/ms_layout/internal/migrator"
)

func main() {
	currentDir, _ := os.Getwd()

	var configPath string
	flag.StringVar(&configPath, "config", config.PathDefault(currentDir), "path to config file")

	var action string
	flag.StringVar(&action, "action", "up", "path to config file")
	flag.Parse()

	var configs *config.Config = configs.MustInit(configPath)

	psqlInfo := buildConnectUrl(configs)
	conn, err := connect_db.New(psqlInfo, configs.Db.Driver)
	if err != nil {
		log.Fatal(err)
	}

	migrator := migrator.MustGetNewMigrator(configs.Db.NameDb)

	switch action {
	case "up":
		if err := migrator.Up(conn); err != nil {
			log.Fatalf("error applying migrations: %v", err)
		} else {
			log.Print("migration has Up successfully")
		}
	case "down":
		if err := migrator.Down(conn); err != nil {
			log.Fatalf("error rolling back migrations: %v", err)
		} else {
			log.Print("migration has Down successfully")
		}
	default:
		log.Fatalf("unknown action: %s. Use 'up' or 'down'.", action)
	}
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
