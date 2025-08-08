package db

import (
	"fmt"
	"github.com/madsilver/task-manager/internal/infrastructure/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewMySQLDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true",
		env.App.DbUser,
		env.App.DbPassword,
		env.App.DbHost,
		env.App.DbPort,
		env.App.DbName,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}
	return conn
}
