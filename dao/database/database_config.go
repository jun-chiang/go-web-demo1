package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	initialOnce sync.Once
)

func init() {
	initialOnce.Do(func() {
		dsn := "root:123456@tcp(127.0.0.1:3306)/go_web_demo1?charset=utf8mb4&parseTime=True&loc=Local"
		tmpDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db = tmpDB
	})
}
