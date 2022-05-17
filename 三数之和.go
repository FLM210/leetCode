// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	i := 0
	for i < len(nums)-2 {
		j, k := i+1, len(nums)-1
		for j != k {
			n1, n2, n3 := nums[j], nums[k], nums[i]
			if n1+n2+n3 == 0 {
				result = append(result, []int{n1, n2, n3})
			} else if n1+n2+n3 > 0 {
				k = k - 1
			} else {
				j = j + 1
			}
		}
		i++
	}
	return result
}

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	res := [][]int{}

// 	for i := 0; i < len(nums)-2; i++ {
// 		n1 := nums[i]
// 		if n1 > 0 {
// 			break
// 		}
// 		if i > 0 && n1 == nums[i-1] {
// 			continue
// 		}
// 		l, r := i+1, len(nums)-1
// 		for l < r {
// 			n2, n3 := nums[l], nums[r]
// 			if n1+n2+n3 == 0 {
// 				res = append(res, []int{n1, n2, n3})
// 				for l < r && nums[l] == n2 {
// 					l++
// 				}
// 				for l < r && nums[r] == n3 {
// 					r--
// 				}
// 			} else if n1+n2+n3 < 0 {
// 				l++
// 			} else {
// 				r--
// 			}
// 		}
// 	}
// 	return res
// }
