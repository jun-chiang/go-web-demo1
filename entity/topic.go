package entity

// 主题结构体
type Topic struct {
	Id         uint64
	Title      string
	Content    string
	CreateTime MyTime
}
