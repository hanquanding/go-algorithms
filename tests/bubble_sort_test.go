package tests

import (
	"fmt"
	"testing"

	"github.com/hqd8080/go-algorithms/algorithms"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{56, 45, 9, 16, 2, 89, 78, 34, 102, 56, 99}
	algorithms.BubbleSort(arr)
	fmt.Println(arr)
}
