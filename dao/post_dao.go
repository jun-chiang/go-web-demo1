package dao

import "github.com/jun-chiang/go-web-demo1/entity"

type PostDao interface {
	QueryPostListByTopicId(id uint) ([]*entity.Post, error)
}
