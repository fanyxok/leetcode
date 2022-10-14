package main

import "fmt"

// 字典序, 例如105
// 1, 10, 100, 101,102,103,104,105, 11,12...
func lexicalOrder(n int) []int {
	soln := []int{}
	for i := 1; i <= 9; i++ {
		if i <= n {
			soln = append(soln, i)
			backtrace(&soln, n, i*10)
		}
	}
	return soln
}

// curNum ending with 0
func backtrace(soln *[]int, n int, curNum int) {
	if curNum > n {
		return
	}
	for i := 0; i <= 9; i++ {
		if curNum+i <= n {
			*soln = append(*soln, curNum+i)
			backtrace(soln, n, (curNum+i)*10)
		} else {
			break
		}
	}
}

func main() {
	fmt.Println(lexicalOrder(123))
}
