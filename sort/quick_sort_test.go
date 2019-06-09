package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Printf("[%s]:%s\n", "快速排序", "排序之前的数据:")
	fmt.Println(arr)

	fmt.Printf("[%s]:%s\n", "快速排序", "排序之后的数据:")
	tmp := QuickSort(arr)
	fmt.Println(tmp)
}
