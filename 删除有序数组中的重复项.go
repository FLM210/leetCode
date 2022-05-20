// 给你一个 升序排列 的数组 nums ,请你 原地 删除重复出现的元素,使每个元素 只出现一次 ,返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

// 由于在某些语言中不能改变数组的长度,所以必须将结果放在数组nums的第一部分。更规范地说,如果在删除重复项之后有 k 个元素,那么nums的前 k 个元素应该保存最终结果。

// 将最终结果插入nums 的前 k 个位置后返回 k 。

// 不要使用额外的空间,你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
package main

import "fmt"

func main() {
	a := []int{1, 1, 1, 2, 2, 3, 4, 5}
	removeDuplicates(a)
}

//遍历数组,遇到重复的值使用分片删除
// func removeDuplicates(nums []int) int {
// 	for n, i := range nums {
// 		for m := n + 1; m < len(nums); {
// 			if i == nums[m] {
// 				nums = append(nums[:m], nums[m+1:]...) //删除重复值
// 				n = 0
// 			} else {
// 				m++
// 			}
// 		}
// 	}
// 	fmt.Println(nums)
// 	return len(nums)
// }

//快慢指针
func removeDuplicates(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j] //当两个指针的值不相等时,将快指针的值前移
		}
	}
	fmt.Println(nums)
	return i + 1
}
