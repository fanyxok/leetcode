package main

func getSum(a int, b int) int {
	all_carry := 1
	for all_carry != 0 {
		sum_no_carry := a ^ b
		all_carry = a & b
		sum_rest := all_carry << 1
		a = sum_no_carry
		b = sum_rest
	}
	return a
}
