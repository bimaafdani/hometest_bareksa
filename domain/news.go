package domain

import (
	"github.com/jinzhu/gorm"
)

// News represent entity of the news
type News struct {
	gorm.Model
	Title   string  `json:"title"`
	Slug    string  `json:"slug"`
	Content string  `json:"content"`
	Status  string  `json:"status"`
	Tags    []Tags  `gorm:"many2many:news_tags;"`
	Topic   []Topic `gorm:"many2many:news_topics;"`
}
