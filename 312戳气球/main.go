package main

// DP

func maxCoins(nums []int) int {
	// coins[i][j] 开区间包括i,不包括j,表示区间i,j戳气球的最大金币
	coins := make([][]int, len(nums)+2)
	for i := range coins {
		coins[i] = make([]int, len(nums)+2)
	}
	// 初始化 coins[i][i] = 0
	// i>j, coins[i][j] = 0
	// i<j, coins[i][j] = max_k{v[i]*v[k]*v[j]+coins[i][k]+coins[k][j]}
	val := make([]int, len(nums)+2)
	val[0], val[len(nums)+1] = 1, 1
	for i := 1; i <= len(nums); i++ {
		val[i] = nums[i-1]
	}
	// 起始位置n-1
	// 结束位置0
	for i := len(nums) - 1; i >= 0; i-- {
		// 起始位置为i+2,包含i和i+1两个元素
		// 结束位置n+1,包含右边界的1
		for j := i + 2; j <= len(nums)+1; j++ {
			// 起始位置i的右边
			// 结束位置j的左边
			for k := i + 1; k < j; k++ {
				another := val[i] * val[j] * val[k]
				sum := coins[i][k] + coins[k][j] + another
				if sum > coins[i][j] {
					coins[i][j] = sum
				}
			}
		}
	}
	return coins[0][len(nums)+1]
}
