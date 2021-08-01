package tests

import (
	"log"
	"testing"

	algorithms "github.com/hqd8080/go-algorithms/algorithms/fibonacci"
)

// 1,1,2,3,5,8,13,21,34,55
func TestFibonacci(t *testing.T) {
	var n int64 = 8
	log.Printf("fib(%d) = %d\n", n, algorithms.Fib(n))

	for i := 1; i <= 10; i++ {
		log.Printf("%d", algorithms.Fib(int64(i)))
	}
}

func TestFibNoRec(t *testing.T) {
	var n int = 8
	log.Printf("fib(%d) = %d\n", n, algorithms.FibNoRec(n))

	for i := 1; i <= 10; i++ {
		log.Printf("%d", algorithms.FibNoRec(i))
	}
}
