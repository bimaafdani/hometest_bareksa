package repository

import "HoteTestBareksa/domain"

// TopicRepository represent repository of the topic
// Expect implementation by the infrastructure layer
type TagsRepository interface {
	Get(id int) (*domain.Tags, error)
	GetAll() ([]domain.Tags, error)
	Save(*domain.Tags) error
	Remove(id int) error
	Update(*domain.Tags) error
}
