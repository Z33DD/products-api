package service

import (
	"eulabs/internal/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Product{})
}
