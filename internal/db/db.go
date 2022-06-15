package db

import (
	"log"

	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {

	var err error

	config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	config.DB.AutoMigrate(
		&models.Allowlist{},
		&models.BlockGroup{},
		&models.Blocklist{},
		&models.BlocklistCategory{},
		&models.Computer{},
		&models.Payment{},
		&models.Settings{},
		&models.User{},
	)
}
