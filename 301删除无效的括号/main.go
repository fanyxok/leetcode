package main

func removeInvalidParentheses(s string) []string {
	l := 0
	r := 0
	for _, v := range s {
		if v == '(' {
			l++
		} else if v == ')' {
			if l == 0 {
				r++
			} else {
				l--
			}
		}
	}
	soln := []string{}
	backtrace(&soln, s, 0, l, r)
	return soln
}

func backtrace(soln *[]string, s string, depth int, l, r int) {
	if r == 0 && l == 0 {
		if isValid(s) {
			*soln = append(*soln, s)
		}
		return
	}

	for i := depth; i < len(s); i++ {
		// 多个c连在一起,删除任何一个都是一样的,否则会重复
		if i != depth && s[i] == s[i-1] {
			continue
		}
		//
		if l+r > len(s)-i {
			return
		}
		//
		if l > 0 && s[i] == '(' {
			backtrace(soln, s[:i]+s[i+1:], i, l-1, r)
		}
		if r > 0 && s[i] == ')' {
			backtrace(soln, s[:i]+s[i+1:], i, l, r-1)
		}
	}
}

func isValid(s string) bool {
	l := 0
	for _, v := range s {
		if v == '(' {
			l++
		} else if v == ')' {
			if l == 0 {
				return false
			} else {
				l--
			}
		}
	}
	return l == 0
}
