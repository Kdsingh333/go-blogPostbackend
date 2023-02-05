package database

import (
	"blog-website/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	} else {
		log.Println("Connect successfully")
	}

	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	   
	)
}
