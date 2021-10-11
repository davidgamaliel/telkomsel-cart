package repository

import (
	"go-cart-api/model"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@/"+os.Getenv("MYSQL_DATABASE")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error connecting to Database")
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.CartItem{}, &model.Transaction{})
	return db

}
