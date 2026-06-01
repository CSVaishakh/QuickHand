package src

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New (url string) (*gorm.DB, error) {
	connection, err := gorm.Open(
		postgres.Open(url),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err;
	}

	return connection,nil
}