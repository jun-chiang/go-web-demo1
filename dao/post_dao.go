package dao

import "github.com/jun-chiang/go-web-demo1/entity"

type PostDao interface {
	// 根据主题Id来查询回复列表
	QueryPostListByTopicId(id uint) ([]*entity.Post, error)
}
