package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Comment      string        `gorm:"type:varchar(255);not null"`
	UserID       uint          `json:"user_id"`
	User         User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArticleLikes []ArticleLike `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
