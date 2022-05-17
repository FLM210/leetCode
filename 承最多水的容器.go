package main

import (
	"fmt"
)

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area, tmp := 0, 0

	for i != j {
		if height[i] < height[j] {
			tmp = (j - i) * height[i]
			i++
		} else {
			tmp = (j - i) * height[j]
			j--
		}
		if area < tmp {
			area = tmp
		}
	}
	return area
}
