package db

import (
	"log"
	"os"

	"github.com/NonNtp/gin-gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("cannot connect to database")
	}

	Conn = db
}

func Migrate() {
	Conn.AutoMigrate(
		&models.User{},
	)
}
