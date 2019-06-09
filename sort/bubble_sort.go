package main

import "fmt"

func main() {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Printf("[%s]:%s\n", "冒泡排序", "排序之前的数据:")
	fmt.Println(arr)

	fmt.Printf("[%s]:%s\n", "冒泡排序", "排序之后的数据:")
	BubbleSort(arr)
	fmt.Println(arr)

}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// 左元素 > 右元素
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
