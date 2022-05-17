// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

package main

import "fmt"

func main() {
	fmt.Println(climbStairs(100))
}
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	f1, f2, f3, i := 1, 2, 3, 3
	for i < n {
		f1, f2 = f2, f3
		f3 = f1 + f2
		i++
	}
	return f3
}
