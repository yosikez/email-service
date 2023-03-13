package database

import (
	"fmt"

	"github.com/yosikez/email-notification-services/config"
	"github.com/yosikez/email-notification-services/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dbConfig, err := config.LoadDatabase()

	if err != nil {
		return fmt.Errorf("failed to load database config : %v", err)
	}

	dsn := dbConfig.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database : %v", err)
	}

	DB = db

	if err := migrate(); err != nil {
		return fmt.Errorf("failed to migrate database : %v", err)
	}

	return nil

}

func migrate() error {
	if err := DB.AutoMigrate(&model.Mail{}); err != nil {
		return err
	}

	return nil
}
