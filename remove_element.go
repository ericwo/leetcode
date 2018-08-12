package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	length := len(nums)
	for _, num := range nums {
		if num == val {
			length--
		}
	}
	return length
}

func main() {
	nums := []int{3, 2, 2, 4}
	fmt.Printf("%v", removeElement(nums, 3))
}
