package main

import (
	"fmt"
)

func main() {
	result := twoSum([]int{1, 2, 3, 40, 5, 60}, 7)
	fmt.Println(result)
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	result := []int{}
	for i := 0; i < len(nums); i++ {
		attended, ok := m[nums[i]]
		if ok {
			result = []int{attended, i}
		} else {
			m[target-nums[i]] = i
		}
	}
	return result
}
