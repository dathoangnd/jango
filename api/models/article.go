package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string `valid:"required" json:"title" gorm:"type:varchar(255);not null"`
	Avatar string `json:"avatar" gorm:"type:varchar(255)"`
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;not null"`
	Categories string `json:"categories" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(500)"`
	Content string `json:"content" gorm:"type:varchar(5000)"`
	Draft string `json:"draft" gorm:"type:varchar(5000)"`
	Publish int `json:"publish" gorm:"type:int;default:0"`
	Slug string `json:"slug" gorm:"type:varchar(255);not null"`
}

type ArticleComment struct {
	gorm.Model
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;not null"`
	ArticleID uint `valid:"int,required" json:"articleId" gorm:"type:int;not null"`
	Content string `valid:"required" json:"content" gorm:"type:varchar(500);not null"`
}

type ArticleLike struct {
	gorm.Model
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;unique_index:user_article;not null"`
	ArticleID uint `valid:"int,required" json:"articleId" gorm:"type:int;unique_index:user_article;not null"`
}