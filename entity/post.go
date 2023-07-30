package entity

import "gorm.io/gorm"

// 主题下面的回复
type Post struct {
	gorm.Model
	ParentId   uint
	Content    string
	CreateTime MyTime
}
