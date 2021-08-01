/**
 * @Author: hqd
 * @Date: 2021/8/1 21:32
 * @Description: consumer
 */

package queue

import "fmt"

// 消费者
func Consumer(data chan int, done chan struct{}) {
	for value := range data {
		fmt.Println("recv:", value)
	}

	// 通知主程序消费结束
	done <- struct{}{}
}
