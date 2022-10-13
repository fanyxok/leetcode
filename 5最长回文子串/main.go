package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	left, right := 0, 0
	for i := range s {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > right-left {
			right = right1
			left = left1
		}
		if right2-left2 > right-left {
			right = right2
			left = left2
		}

	}
	return s[left : right+1]
}

func expandAroundCenter(s string, i1, i2 int) (int, int) {
	for ; i1 >= 0 && i2 < len(s) && s[i1] == s[i2]; i1, i2 = i1-1, i2+1 {
	}
	return i1 + 1, i2 - 1
}
