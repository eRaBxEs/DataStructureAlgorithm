package main

import (
	"fmt"
)

func allEven(nums []int) bool {
	var isEven bool
	var evenCount int
	var oddCount int

	for _, num := range nums {
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	if evenCount == len(nums) {
		isEven = true
	} else {
		isEven = false
	}

	return isEven
}

func main() {

	fmt.Println(allEven([]int{4, 90, 68, 6, -2}))
	fmt.Println(allEven([]int{14, 40, 36, 3}))
	fmt.Println(allEven([]int{30, 24, 2048, 0, 12, 50}))
	fmt.Println(allEven([]int{7, 7, 7, 7}))
	fmt.Println(allEven([]int{100}))
	fmt.Println(allEven([]int{1, 2, 4, 6, 8}))
	fmt.Println(allEven([]int{42, 18, 96, 4, 70, 12, 58, 30, 84, 26}))
	fmt.Println(allEven([]int{-5, -3, -4}))

}
