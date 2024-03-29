/**
 * @Author: hqd
 * @Date: 2021/8/1 20:14
 * @Description: 快速排序
 */

package algorithms

func QuickSort(arr []int) []int {
	n := len(arr)

	// 递归结束条件
	if n <= 1 {
		temp := make([]int, n)
		copy(temp, arr)
		return temp
	}

	// 使用第一个元素作为基准值
	pivot := arr[0]

	// 小元素和大元素各分成一个数组
	low := make([]int, 0, n)
	high := make([]int, 0, n)

	// 分治
	// 小的元素放 low[]
	// 大的元素放 high[]
	for i := 1; i < n; i++ {
		if arr[i] < pivot {
			low = append(low, arr[i])
		} else {
			high = append(high, arr[i])
		}
	}

	// 子区间递归快排，分治排序
	low, high = QuickSort(low), QuickSort(high)

	return append(append(low, pivot), high...)
}
