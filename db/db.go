package db

import (
	"github.com/akshaybt001/DatingApp_Payment_Service/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(connectTo string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Payment{})
	db.AutoMigrate(&entities.Subscription{})
	db.AutoMigrate(&entities.UserSubscription{})
	return db, nil
}
