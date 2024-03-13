package config

import (
	"fiber-go/model"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Models = []interface{}{
	&model.Cashier{},
	&model.Category{},
	&model.Payment{},
	&model.PaymentType{},
	&model.Product{},
	&model.Discount{},
	&model.Order{},
}

func AutoMigrate(connection *gorm.DB) {
	err := connection.Debug().AutoMigrate(Models...)
	if err != nil {
		color.Red(err.Error())
	}
}

func Connect() {
	godotenv.Load()
	dbhost := os.Getenv("MYSQL_HOST")
	dbport := os.Getenv("MYSQL_PORT")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)
	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

	fmt.Println("Database connected!")

	AutoMigrate(db)
}
