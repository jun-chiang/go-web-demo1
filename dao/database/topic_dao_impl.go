package database

import (
	"github.com/jun-chiang/go-web-demo1/entity"
)

var (
	topicDaoImpl *TopicDaoImpl
)

type TopicDaoImpl struct {
}

func NewTopicDaoImplInstance() *TopicDaoImpl {
	initialOnce.Do(func() {
		topicDaoImpl = &TopicDaoImpl{}
	})
	return topicDaoImpl
}

// 实现Topic_dao的方法
func (*TopicDaoImpl) QueryTopicById(id uint) (*entity.Topic, error) {
	topic := &entity.Topic{}
	db.Find(&topic, id)
	// TODO 如何获取查询过程中出现的异常呢。
	return topic, nil
}
