package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"week1/global"
)

func initDB() {
	dsn := AppConfig.DataBase.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour * 1)
	sqlDB.SetMaxOpenConns(AppConfig.DataBase.MaxOpenConns)
	sqlDB.SetMaxIdleConns(AppConfig.DataBase.MaxIdleConns)

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	global.Db = db
}
