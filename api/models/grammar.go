package models

import "github.com/jinzhu/gorm"

type GrammarLike struct {
	gorm.Model
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;unique_index:user_grammar;not null"`
	GrammarOrder string `valid:"required" json:"grammarOrder" gorm:"type:varchar(6);unique_index:user_grammar;not null"`
}