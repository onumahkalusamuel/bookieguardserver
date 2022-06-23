package db

import (
	"log"
	"os"

	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {

	var err error

	if config.SERVER_HOST == "localhost" {
		config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	} else {
		dsn := os.Getenv("DATABASE_URL")
		config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatalln(err)
	}

	config.DB.AutoMigrate(
		&models.Allowlist{},
		&models.BlockGroup{},
		&models.Blocklist{},
		&models.BlocklistCategory{},
		&models.Computer{},
		&models.Host{},
		&models.Payment{},
		&models.Settings{},
		&models.User{},
	)
}
