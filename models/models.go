package models

import (
	"fmt"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg config.Config) error {
	var err error
	db, err = connectDB(cfg)
	if err != nil {
		return err
	}

	return autoMigrates()
}

func GetterDB() *gorm.DB {
	return db
}

func connectDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBConfig.User, cfg.DBConfig.Password, cfg.DBConfig.Host, cfg.DBConfig.DataBaseName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func autoMigrates() error {
	return db.AutoMigrate(&Reservation{})
}
