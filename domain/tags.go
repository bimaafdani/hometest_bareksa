package domain

import (
	"github.com/jinzhu/gorm"
)

type Tags struct {
	gorm.Model
	Tags string `json:"name_tags"`
	News []News `gorm:"many2many:news_tags;"`
}
