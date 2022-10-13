package main

import (
	"sort"
)

func countSmaller(nums []int) []int {
	N := len(nums)
	soln := make([]int, N)
	sorted := []int{}
	for i := N - 1; i >= 0; i-- {
		idx := sort.Search(len(sorted), func(x int) bool { return sorted[x] >= nums[i] })
		sorted = append(sorted[:idx], append([]int{nums[i]}, sorted[idx:]...)...)
		soln[i] = idx
	}
	return soln
}
