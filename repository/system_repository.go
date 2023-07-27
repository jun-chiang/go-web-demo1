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
	TopicIndexMap map[uint64]*Topic
	// 定义并初始化话题回复索引
	PostIndexMap map[uint64][]*Post = make(map[uint64][]*entity.Post, 5)
)

// 从文件初始话题索引
func initTopicIndexMap() error {
	// 读取json文件
	jsonFile, err := os.Open("initial_data.json")
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
	json.Unmarshal(byteValue, &topics)
	// 初始化话题索引变量
	TopicIndexMap = make(map[uint64]*Topic, 5)
	for i := range topics {
		topic := topics[i]
		// 以ID为key，把对象指针放到Map里面去
		TopicIndexMap[topic.Id] = &topic
	}
	return nil
}
