//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

package main

import "fmt"

func main() {
	nums := []int{1, 2}
	rotate(nums, 3)
	fmt.Println(nums)

}

func rotate(nums []int, k int) {
	k = k % len(nums)
	fmt.Println(k)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
