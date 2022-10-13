package main

import "math"

func lengthOfLongestSubstring(s string) int {
	pos := make([]int, 128) //记录字符上一次出现位置
	for i := range pos {
		pos[i] = -1
	}
	maxsub := 0
	start := 0 // 滑动窗口起始位置
	for i, v := range s {
		start = int(math.Max(float64(start), float64(pos[v]+1)))    // 遇到重复字符串，并且上一次出现位置在窗口右侧的，更新窗口
		maxsub = int(math.Max(float64(maxsub), float64(i-start+1))) // 比较是老的sub更大，还是新的sub更大
		pos[v] = i
	}
	return maxsub
}
