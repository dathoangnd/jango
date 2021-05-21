package models

import "github.com/jinzhu/gorm"

type Thread struct {
	gorm.Model
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;not null"`
	Categories string `json:"categories" gorm:"type:varchar(255)"`
	Content string `valid:"required" json:"content" gorm:"type:varchar(500);not null"`
	Vote int `valid:"int" json:"vote" gorm:"type:int;default:0"`
}

type ThreadReply struct {
	gorm.Model
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;not null"`
	ThreadId uint `valid:"int,required" json:"threadId" gorm:"type:int;not null"`
	Content string `valid:"required" json:"content" gorm:"type:varchar(500);not null"`
	Vote int `valid:"int" json:"vote" gorm:"type:int;default:0"`
}