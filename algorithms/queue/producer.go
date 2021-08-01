/**
 * @Author: hqd
 * @Date: 2021/8/1 21:32
 * @Description: producer
 */

package queue

func Producer(data chan int) {
	for i := 0; i < 10; i++ {
		// 发送数据
		data <- i
	}
	// 生产结束，关闭通道
	close(data)
}
