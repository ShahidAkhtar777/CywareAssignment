package utils

import (
	"CywareAssignment/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

func InitializeDB() {
	dbUser := viper.GetString("DB_USER")
	dbName := viper.GetString("DB_NAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbSSLMode := viper.GetString("DB_SSL_MODE")

	if dbUser == "" || dbName == "" || dbPassword == "" {
		log.Fatal("Missing database configuration")
	}

	connectionString := "postgres://" + dbUser + ":" + dbPassword + "@/" + dbName + "?sslmode=" + dbSSLMode
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	DB = db

	// Run auto-migration to create tables
	// TODO: Use migrations folder for all migrations.
	DB.AutoMigrate(&models.CronRun{})
}
