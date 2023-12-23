package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "github.com/donnyirianto/go-clean/pkg/config"
	domain "github.com/donnyirianto/go-clean/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	mysqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.User{})

	return db, dbErr
}
