/**
 * @Author: hqd
 * @Description: lru_cache_test
 * @File:  lru_cache_test
 * @Version: 1.0.0
 * @Date: 2021/1/25 17:56
 */

package tests

import (
	"fmt"
	"testing"

	"github.com/hqd8080/go-algorithms/algorithms"
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
