package main

import "fmt"

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	i, j := 1, len(height)
	area := 0
	if height[i-1] < height[j-1] {
		for i != j {
			if height[i-1] < height[j-1] {
				tmp := (j - i) * height[i-1]
				if tmp > area {
					area = tmp
				}
			} else {
				tmp := (j - i) * height[j-1]
				if tmp > area {
					area = tmp
				}
			}
			i++
		}
	} else {
		for j != i {
			if height[i-1] < height[j-1] {
				tmp := (j - i) * height[i-1]
				if tmp > area {
					area = tmp
				}
			} else {
				tmp := (j - i) * height[j-1]
				if tmp > area {
					area = tmp
				}
			}
			j--
		}
	}

	return area
}
