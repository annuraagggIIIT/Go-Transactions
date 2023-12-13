package db

import (
	models "example/v2/model"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect(){
	godotenv.Load()
	dbhost:=os.Getenv("MYSQL_HOST")
	dbuser:=os.Getenv("MYSQL_USER")
	dbpassword:=os.Getenv("MYSQL_PASSWORD")
	dbname:=os.Getenv("MYSQL_DBNAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil{
		fmt.Println("error",err)

	
	}
	DB=db

	AutoMigrate(db)
	fmt.Println("db connected succesfully")
}
func AutoMigrate(connection *gorm.DB){
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Discount{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},

	)
}