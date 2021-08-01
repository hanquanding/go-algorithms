/**
 * @Author: hqd
 * @Date: 2021/8/1 20:01
 * @Description: 冒泡排序
 */

package algorithms

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// 左元素 > 右元素
			if arr[j] > arr[j+1] {
				// 交换，从小到大排序
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
