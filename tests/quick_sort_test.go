package tests

import (
	"fmt"
	"testing"

	"github.com/hqd8080/go-algorithms/algorithms"
)

func TestQuickSort(t *testing.T) {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	tmp := algorithms.QuickSort(arr)
	fmt.Println(tmp)
}
