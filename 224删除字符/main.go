package main

/**
 * @param str: the string
 * @param k: the length
 * @return: the substring with  the smallest lexicographic order
 */
func DeleteChar(str string, k int) string {
	// Write your code here.
	if len(str) <= k {
		return str
	}
	soln := str[len(str)-k:]
	for i := len(str) - k - 1; i >= 0; i-- {
		if str[i] <= soln[0] {
			soln = update(soln, str[i])
		}
	}
	return soln
}

func update(soln string, b byte) string {
	upd := false
	for i := 1; i < len(soln); i++ {
		if soln[i-1] > soln[i] {
			soln = string(b) + soln[:i-1] + soln[i:]
			upd = true
			break
		}
	}
	if !upd {
		soln = string(b) + soln[:len(soln)-1]
	}
	return soln
}
