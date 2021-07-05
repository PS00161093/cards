package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var indexes = []int{0, 0}
	for a, i := range nums {
		for b, j := range nums {
			if i+j == target {
				indexes[0] = a
				indexes[1] = b
				break
			}
		}
	}
	return indexes
}
