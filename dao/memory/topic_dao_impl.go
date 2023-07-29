package memory

import (
	"sync"

	"github.com/jun-chiang/go-web-demo1/entity"
	"github.com/jun-chiang/go-web-demo1/repository"
)

// 定义结构体
type TopicDaoImpl struct {
}

var (
	// 声明结构体变量
	topicDaoImpl *TopicDaoImpl
	initialOnce  sync.Once
)

// 获取TopicDaoImpl的全局唯一实例
func NewTopicDaoImplInstance() *TopicDaoImpl {
	// 单例模式
	initialOnce.Do(func() {
		// 为结构体变量赋值
		topicDaoImpl = &TopicDaoImpl{}
	})
	// 返回已经赋值的topicDaoImpl结构体对象指针
	return topicDaoImpl
}

// 实现Topic_dao的方法
func (*TopicDaoImpl) QueryTopicById(id uint64) (*entity.Topic, error) {
	return repository.TopicIndexMap[id], nil
}
