package db

import (
	"fmt"
	models "github.com/kaspers1778/money-processing-svc/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"))

type DBInstance struct {
	DB *gorm.DB
}

func ConnectDB() DBInstance {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Fail to connect to db: ", err)
	}
	log.Println("Connected to DB.")
	log.Println("Migrating data.")
	err = InitialMigration(db)
	if err != nil {
		log.Fatal("Fail to perform migration: ", err)
	}
	return DBInstance{db}
}

func InitialMigration(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&models.Account{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Client{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		return err
	}
	return nil
}
