package tests

import (
	"fmt"
	"testing"

	algorithms "github.com/hqd8080/go-algorithms/algorithms/sort"
)

func TestQuickSort(t *testing.T) {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	tmp := algorithms.QuickSort(arr)
	fmt.Println(tmp)
}
