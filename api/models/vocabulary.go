package models

import "github.com/jinzhu/gorm"

type Voca struct {
	gorm.Model
	Text string `valid:"string,required" json:"text" gorm:"type:varchar(255);not null"`
	Read string `valid:"string" json:"read" gorm:"type:varchar(255)"`
	Meaning string `valid:"string,required" json:"meaning" gorm:"type:varchar(255);not null"`
	ExampleText string `valid:"string,required" json:"exampleText" gorm:"type:varchar(500);not null"`
	ExampleRead string `valid:"string" json:"exampleRead" gorm:"type:varchar(500)"`
	ExampleMeaning string `valid:"string,required" json:"exampleMeaning" gorm:"type:varchar(500);not null"`
	Lesson int `valid:"int,required" json:"lesson" gorm:"type:int;not null"`
}

type VocaLike struct {
	gorm.Model
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;unique_index:user_voca;not null"`
	VocaOrder uint `valid:"int,required" json:"vocaOrder" gorm:"type:int;unique_index:user_voca;not null"`
}