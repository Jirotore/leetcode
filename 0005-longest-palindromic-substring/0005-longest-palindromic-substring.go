func longestPalindrome(s string) string {
	srune := []rune(s)
    if len(srune) < 2 {
        return s
    }
	var longest []rune
    for i := 0; i < len(srune); i++ {
		for j := i; j < len(srune); j++ {
            buf := srune[i:j+1]
			if !check(buf) {
				continue
			}
			if len(longest) < len(buf) {
				longest = buf
			}
		}
	}
	return string(longest)
}

func check(s []rune) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
