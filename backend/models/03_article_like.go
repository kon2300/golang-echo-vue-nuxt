package models

import "gorm.io/gorm"

type ArticleLike struct {
	gorm.Model
	ArticleID uint    `json:"article_id" gorm:"index:idx_name,unique"`
	UserID    uint    `json:"user_id" gorm:"index:idx_name,unique"`
	Article   Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
