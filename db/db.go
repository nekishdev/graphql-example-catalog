package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/nekishdev/graphql-example-catalog/gorm_models"
	"os"
	"time"
)

var db *gorm.DB // DB instance

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Создать строку подключения
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(fmt.Sprintf("Error open connection to DB. %s", err))
	}

	conn.DB().SetConnMaxLifetime(time.Hour)
	conn.DB().SetMaxOpenConns(100)
	conn.DB().SetMaxIdleConns(10)
	db = conn

	db.Debug().AutoMigrate(
		&gorm_models.Product{},
		&gorm_models.Category{},
		&gorm_models.Property{},
		&gorm_models.ProductProperty{},
	)
}

func GetDB() *gorm.DB {
	return db
}
