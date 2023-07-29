package dao

import (
	"github.com/jun-chiang/go-web-demo1/entity"
	"github.com/jun-chiang/go-web-demo1/repository"
)

// 定义结构体
type PostDaoImpl struct {
}

var postDaoImpl *PostDaoImpl

// 获取TopicDaoImpl的全局唯一实例
func NewPostDaoImplInstance() *PostDaoImpl {
	// 单例模式
	initialOnce.Do(func() {
		// 为结构体变量赋值
		postDaoImpl = &PostDaoImpl{}
	})
	// 返回已经赋值的topicDaoImpl结构体对象指针
	return postDaoImpl
}

// 实现Topic_dao的方法
func (*PostDaoImpl) QueryPostListByTopicId(id uint64) ([]*entity.Post, error) {
	return repository.PostIndexMap[id], nil
}
