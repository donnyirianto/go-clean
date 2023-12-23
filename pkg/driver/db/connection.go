package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "github.com/donnyirianto/go-clean/pkg/config"
	domain "github.com/donnyirianto/go-clean/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=latin1&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, dbErr := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.User{})

	return db, dbErr
}
