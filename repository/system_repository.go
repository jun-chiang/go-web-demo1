package repository

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jun-chiang/go-web-demo1/entity"
)

// 为类型定义别名 如果没有 ‘=’ 就是定义新的类型
type Topic = entity.Topic
type Post = entity.Post

// 存储数据的变量
var (
	// 定义话题索引
	TopicIndexMap map[uint]*Topic = make(map[uint]*Topic, 5)
	// 定义并初始化话题回复索引
	PostIndexMap map[uint][]*Post = make(map[uint][]*entity.Post, 5)
)

// 从文件初始话题索引
func InitTopicIndexMap() error {
	// 读取json文件
	jsonFile, err := os.Open("repository/initial_data.json")
	if err != nil {
		return err
	}
	// 延迟关闭文件流
	defer jsonFile.Close()
	// 读取全部文件
	var topics []Topic
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	// json反序列化
	err = json.Unmarshal(byteValue, &topics)
	if err != nil {
		return err
	}
	for i := range topics {
		topic := topics[i]
		// 以ID为key，把对象指针放到Map里面去
		TopicIndexMap[topic.ID] = &topic
	}
	return nil
}
