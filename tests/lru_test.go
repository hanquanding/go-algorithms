package tests

import (
	"fmt"
	"testing"

	algorithms "github.com/hqd8080/go-algorithms/algorithms/lru"
)

func TestLRUCache(t *testing.T) {
	cache := algorithms.NewLRUCache(2)
	cache.Add("aa", 1)
	cache.Add("bb", 2)
	cache.Add("cc", 3)

	fmt.Println(cache.Get("aa"))
	fmt.Println(cache.Get("bb"))
	fmt.Println(cache.Get("cc"))
}
