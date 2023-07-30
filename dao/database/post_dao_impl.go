package database

import "github.com/jun-chiang/go-web-demo1/entity"

type PostDaoImpl struct{}

var postDaoImpl *PostDaoImpl

func NewPostDaoImplInstance() *PostDaoImpl {
	initialOnce.Do(func() {
		postDaoImpl = &PostDaoImpl{}
	})
	return postDaoImpl
}

// 实现postDao的方法
func (*PostDaoImpl) QueryPostListByTopicId(id uint) ([]*entity.Post, error) {
	var postList []*entity.Post
	db.Where(&entity.Post{
		ParentId: id,
	}).Find(&postList)
	return postList, nil
}
