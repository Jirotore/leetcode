func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxlen := 0, 0
	check := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		current := right - left - 1
		if current > maxlen {
			maxlen = current
			start = left + 1
		}
	}

	for i := 0; i < len(s); i++ {
		check(i, i)
		check(i, i+1)
	}

	return s[start : start+maxlen]
}
