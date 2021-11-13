package persistence

import (
	"HoteTestBareksa/domain"
	"HoteTestBareksa/domain/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewsRepositoryImpl Implements repository.NewsRepository
type TagsRepositoryImpl struct {
	Conn *gorm.DB
}

// NewNewsRepositoryWithRDB returns initialized NewsRepositoryImpl
func NewTagsRepositoryWithRDB(conn *gorm.DB) repository.TagsRepository {
	return &TagsRepositoryImpl{Conn: conn}
}

func (r *TagsRepositoryImpl) Save(tags *domain.Tags) error {
	if err := r.Conn.Save(&tags).Error; err != nil {
		return err
	}
	return nil
}

// Get topic by id return domain.topic
func (r *TagsRepositoryImpl) Get(id int) (*domain.Tags, error) {
	tag := &domain.Tags{}
	if err := r.Conn.Preload("News").First(&tag, id).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *TagsRepositoryImpl) GetAll() ([]domain.Tags, error) {
	tags := []domain.Tags{}
	if err := r.Conn.Preload("News").Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// Remove delete topic
func (r *TagsRepositoryImpl) Remove(id int) error {
	tag := &domain.Tags{}
	if err := r.Conn.First(&tag, id).Error; err != nil {
		return err
	}

	if err := r.Conn.Delete(&tag).Error; err != nil {
		return err
	}

	return nil
}

// Update data topic
func (r *TagsRepositoryImpl) Update(tag *domain.Tags) error {
	if err := r.Conn.Model(&tag).UpdateColumns(domain.Tags{Tags: tag.Tags}).Error; err != nil {
		return err
	}

	return nil
}
