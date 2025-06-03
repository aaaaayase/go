package config

import (
	"first-project/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
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
