package entity

import "gorm.io/gorm"

// 主题结构体
type Topic struct {
	gorm.Model
	Title      string
	Content    string
	CreateTime MyTime
}
