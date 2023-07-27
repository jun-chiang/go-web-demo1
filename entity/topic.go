package entity

import "time"

// 主题结构体
type Topic struct {
	Id         uint64
	Title      string
	Content    string
	CreateTime time.Time
}
