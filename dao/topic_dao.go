package dao

import "github.com/jun-chiang/go-web-demo1/entity"

// TopicDao定义了查询Topic的方式
type TopicDao interface {
	// 根据Topic的ID查询Topi对象
	QueryTopicById(id uint64) (*entity.Topic, error)
}
