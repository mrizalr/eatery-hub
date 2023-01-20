package api

import (
	"log"

	"github.com/mrizalr/eatery-hub/config"
	"github.com/mrizalr/eatery-hub/internal/models"
	"github.com/mrizalr/eatery-hub/internal/server"
	"github.com/mrizalr/eatery-hub/pkg/db/mysql"
)

func StartApplication() {
	configFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	config, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	db, err := mysql.NewMysqlDB(config)
	if err != nil {
		log.Fatalf("MysqlConnect: %v", err)
	} else {
		db.AutoMigrate(&models.User{})
		log.Print("Mysql connected")
	}

	server := server.New(config, db)
	if err := server.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
