package main

import (
	"fmt"
	"sort"
)

// 暴力超时
type Slice struct {
	Q int
	W int
	E int
}

func threeSum(nums []int) [][]int {
	num := sort.IntSlice(nums)
	sort.Sort(num)
	soln := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if num[i] > 0 {
			break
		}
		// 去重
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if v := num[i] + num[l] + num[r]; v == 0 {
				soln = append(soln, []int{nums[i], nums[l], nums[r]})
				// 处理重复
				for l < r && num[l] == num[l+1] {
					l++
				}
				for l < r && num[r] == num[r-1] {
					r--
				}
				l++
				r--
			} else if v > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return soln
}

func main() {
	for _, v := range threeSum([]int{-1, 0, 1, 2, -1, -4}) {
		fmt.Println(v)
	}
}
