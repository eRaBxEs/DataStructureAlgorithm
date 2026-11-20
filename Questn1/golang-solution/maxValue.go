package main

import (
	"fmt"
)

func maxValue(nums []int) int {
	max := nums[0]

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}

func main() {
	fmt.Println(maxValue([]int{4, 7, 2, 8, 10, 9}))
	fmt.Println(maxValue([]int{10, 5, 40, 41}))
	fmt.Println(maxValue([]int{-5, -2, -1, -11}))
	fmt.Println(maxValue([]int{42}))
	fmt.Println(maxValue([]int{1000, 8}))
	fmt.Println(maxValue([]int{1000, 8, 9000}))
	fmt.Println(maxValue([]int{2, 5, 1, 1, 4}))
}
