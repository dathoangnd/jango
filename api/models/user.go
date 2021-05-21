package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `valid:"email,required" json:"email" gorm:"type:varchar(255);unique_index;unique;not null"`
	Password string `valid:"required" json:"-" gorm:"type:varchar(255);not null"`
	Name string `valid:"required" json:"name" gorm:"type:varchar(50);not null"`
	Avatar string `json:"avatar" gorm:"type:varchar(255)"`
	Birthday int64 `valid:"int,optional" json:"birthday" gorm:"type:bigint"`
	Gender int `json:"gender" gorm:"type:int"`
  City string `json:"city" gorm:"varchar(50)"`
	FacebookId string `json:"facebookId" gorm:" varchar(255)"`
	About string `json:"about" gorm:"varchar(500)"`
 }