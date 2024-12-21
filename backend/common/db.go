package common

import (
	"fmt"
	"log"

	"github.com/nandonyata/kerjaaja/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var tables = []interface{}{
	model.Role{},
	model.User{},
}

func ConnectDB(logger *logrus.Logger) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		Get().DB.Host,
		Get().DB.User,
		Get().DB.Password,
		Get().DB.Schema,
		Get().DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database. error: %v", err)
		return err
	}

	DB = db

	if Get().DB.MigrateSchema {
		for _, t := range tables {
			if err := db.AutoMigrate(&t); err != nil {
				logger.Fatalf("error migrating schema. error: %v", err)
				return err
			}
		}
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
