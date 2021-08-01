/**
 * @Author: hqd
 * @Date: 2021/8/1 20:37
 * @Description: 斐波那契数列
 */

package fibonacci

// 递归函数实现斐波那契数列
func Fib(n int64) int64 {
	if n <= 1 {
		return n
	}

	return Fib(n-2) + Fib(n-1)
}

// 使用循环实现斐波那契数列
func FibNoRec(n int) int {
	var f1 int = 0
	var f2 int = 1

	for i := 0; i < n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f1
}

// 递归算法的时间复杂度为：O（2^n）2的N次方，随着N的增大，呈现指数增长
// 非递归算法的时间复杂度为：O（n）O的N次方
