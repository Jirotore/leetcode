package leetcode

func findAnagrams(s string, p string) []int {
	res := make([]int, 0, len(s)/len(p))

	m := map[SignWordType]struct{}{signWord(p): struct{}{}}

	for i := 0; i <= len(s)-len(p); i++ {
		if _, ok := m[signWord(s[i:i+len(p)])]; ok {
			res = append(res, i)
		}
	}

	return res
}

type SignWordType [26]int

func signWord(s string) SignWordType {
	m := SignWordType{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := range s {
		m[s[i]-'a']++
	}

	return SignWordType(m)
}
