package application

import (
	"HoteTestBareksa/config"
	"HoteTestBareksa/domain"

	"HoteTestBareksa/infrastructure/persistence"
)

// GetTopic returns a topic by id
func GetTag(id int) (*domain.Tags, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewTagsRepositoryWithRDB(conn)
	return repo.Get(id)
}

// GetAllTopic return all topics
func GetAllTag() ([]domain.Tags, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewTagsRepositoryWithRDB(conn)
	return repo.GetAll()
}

func SaveTags(name_tags string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewTagsRepositoryWithRDB(conn)
	u := &domain.Tags{
		Tags: name_tags,
	}
	return repo.Save(u)
}

// RemoveTopic do remove topic by id
func RemoveTag(id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewTagsRepositoryWithRDB(conn)
	return repo.Remove(id)
}

// UpdateTopic do update topic by id
func UpdateTag(p domain.Tags, id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewTagsRepositoryWithRDB(conn)
	p.ID = uint(id)

	return repo.Update(&p)
}
