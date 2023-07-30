package service

import (
	"errors"
	"strconv"
	"sync"

	"github.com/jun-chiang/go-web-demo1/dao"
	"github.com/jun-chiang/go-web-demo1/dao/database"
	"github.com/jun-chiang/go-web-demo1/dao/memory"
	"github.com/jun-chiang/go-web-demo1/entity"
)

type QueryPageInfoFlow struct {
	topicIdStr string
	topicId    uint
	PageInfo   *entity.PageInfo
	topic      *entity.Topic
	postList   []*entity.Post
}

// 通过TopicId查询页面信息的方法
func QueryPageInfo(topicIdStr string) (*entity.PageInfo, error) {
	return NewQueryPageInfoFlow(topicIdStr).Do()
}

// 工厂函数，返回查询流程结构体的指针，用于启动流程
func NewQueryPageInfoFlow(topicIdStr string) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicIdStr: topicIdStr,
	}
}

// 查询页面信息的整个流程
func (f *QueryPageInfoFlow) Do() (*entity.PageInfo, error) {
	// 参数校验
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	// 准备数据
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	// 组装实体
	if err := f.packageInfo(); err != nil {
		return nil, err
	}
	return f.PageInfo, nil
}

// 参数校验
func (f *QueryPageInfoFlow) checkParam() error {
	topicId, err := strconv.ParseUint(f.topicIdStr, 0, 64)
	if err != nil {
		return errors.New("parse topicIdStr to uint64 failed")
	}
	if topicId <= 0 {
		return errors.New("topic id must larger than 0")
	}
	f.topicId = uint(topicId)
	return nil
}

// 获取Topic的信息以及对应的Post列表的信息
func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	var topicErr, postErr error
	// 获取Topic的信息
	go func() {
		defer wg.Done()
		var topicDao dao.TopicDao = database.NewTopicDaoImplInstance()
		topic, err := topicDao.QueryTopicById(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.topic = topic
	}()
	// 获取Post列表的信息
	go func() {
		defer wg.Done()
		// 利用接口实现多态
		var postDao dao.PostDao = memory.NewPostDaoImplInstance()
		postList, err := postDao.QueryPostListByTopicId(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.postList = postList
	}()
	wg.Wait()

	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}
	return nil
}

// 包装数据
func (f *QueryPageInfoFlow) packageInfo() error {
	f.PageInfo = &entity.PageInfo{
		Topic:    f.topic,
		PostList: f.postList,
	}
	return nil
}
