package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null;unique"`
	Articles []Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
