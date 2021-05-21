package models

import "github.com/jinzhu/gorm"

type KanjiLike struct {
	gorm.Model
	UserID uint `valid:"int,required" json:"userId" gorm:"type:int;unique_index:user_kanji;not null"`
	KanjiOrder string `valid:"required" json:"kanjiOrder" gorm:"type:varchar(7);unique_index:user_kanji;not null"`
}