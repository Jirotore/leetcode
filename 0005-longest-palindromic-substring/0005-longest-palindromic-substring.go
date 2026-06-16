func longestPalindrome(s string) string {
	// srune := []rune(s)
    if len(s) < 2 {
        return s
    }
	var longest string
    for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
            buf := s[i:j+1]
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

func check(s string) bool {
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
