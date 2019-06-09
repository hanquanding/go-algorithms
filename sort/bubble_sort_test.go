package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{56, 45, 9, 16, 2, 89, 78, 34, 102, 56, 99}
	fmt.Printf("[%s]:%s\n", "冒泡排序", "排序之前的数据:")
	fmt.Println(arr)

	fmt.Printf("[%s]:%s\n", "冒泡排序", "排序之后的数据:")
	BubbleSort(arr)
	fmt.Println(arr)
}
