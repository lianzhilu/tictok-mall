package database

import (
	"github.com/lianzhilu/tiktokmall/pkg/store/model"
	"gorm.io/gorm"
)

var models = []interface{}{
	model.User{},
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}
