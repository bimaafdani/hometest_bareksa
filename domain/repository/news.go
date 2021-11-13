package repository

import "HoteTestBareksa/domain"

type NewsRepository interface {
	Get(id int) (*domain.News, error)
	GetAll() ([]domain.News, error)
	GetBySlug(slug string) ([]*domain.News, error)
	GetAllByStatus(status string) ([]domain.News, error)
	Save(*domain.News) error
	Remove(id int) error
	Update(*domain.News) error
}
