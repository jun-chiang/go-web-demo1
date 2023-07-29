package dao

import "github.com/jun-chiang/go-web-demo1/entity"

type PostDao interface {
	QueryPostListByTopicId(id uint64) ([]*entity.Post, error)
}
