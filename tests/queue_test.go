package tests

import (
	"testing"

	algorithms "github.com/hqd8080/go-algorithms/algorithms/queue"
)

func TestQueue(t *testing.T) {
	done := make(chan struct{})        // 用于接收消费者结束信号
	data := make(chan int)             // 数据管道
	go algorithms.Consumer(data, done) // 启动消费者
	go algorithms.Producer(data)       // 启动生产者
	<-done                             // 阻塞，直到消费者发回结束信号
}
