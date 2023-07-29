package entity

// 主题下面的回复
type Post struct {
	Id         uint64
	ParentId   uint64
	Content    string
	CreateTime MyTime
}
