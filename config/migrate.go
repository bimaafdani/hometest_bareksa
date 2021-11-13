package config

import (
	"HoteTestBareksa/domain"
	"log"

	"github.com/jinzhu/gorm"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.AutoMigrate(domain.News{}, domain.Topic{}, domain.Tags{})
	log.Println("Migration has been processed")

	return conn, nil
}
