package models

import (
	"os"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDatabase() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(
		&User{},
		&Article{},
		&ArticleComment{},
		&ArticleLike{},
		&Voca{},
		&VocaLike{},
		&GrammarLike{},
		&KanjiLike{},
		&Document{},
		&Thread{},
		&ThreadReply{},
	)

	log.Println("Database connected")

	return db
}

func GetDb() *gorm.DB {
	return db
}