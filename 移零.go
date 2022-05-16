//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
package main

import "fmt"

func main() {
	t := []int{1, 2, 0, 3, 0, 4, 5}
	moveZeroes(t)
	fmt.Println(t)

}

func moveZeroes(nums []int) {
	count := 0
	for _, m := range nums {
		if m != 0 {
			nums[count] = m
			count++
		}
	}
	for count < len(nums) {
		nums[count] = 0
		count++
	}
}
