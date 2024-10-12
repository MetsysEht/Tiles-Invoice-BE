package database

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/database/models"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/gormDatabase"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/logger"
	"gorm.io/gorm"
)

func GetDatabase(c *gormDatabase.Config) (*gorm.DB, error) {
	db, err := gormDatabase.CreateGormDatabase(c)
	if err != nil {
		logger.L.Fatal("Could not connect to DB")
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
