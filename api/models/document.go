package models

import "github.com/jinzhu/gorm"

type Document struct {
	gorm.Model
	Title string `valid:"required" json:"title" gorm:"type:varchar(255);not null"`
	Avatar string `json:"avatar" gorm:"type:varchar(255)"`
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;not null"`
	Categories string `json:"categories" gorm:"type:varchar(255)"`
	Content string `json:"content" gorm:"type:varchar(5000)"`
	Publish int `json:"publish" gorm:"type:int;default:0"`
	Slug string `valid:"required" json:"slug" gorm:"type:varchar(255);not null"`
	File string `valid:"required" json:"file" gorm:"type:varchar(255)"`
	Author string `json:"author" gorm:"type:varchar(255)"`
	Source string `json:"source" gorm:"type:varchar(255)"`
	Download uint `valid:"int" json:"download" gorm:"type:int;default:0"`
}