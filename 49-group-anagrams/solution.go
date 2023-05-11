package leetcode

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0, len(strs))

	hashMap := make(map[SignWordType][]string, len(strs))

	for i := 0; i < len(strs); i++ {
		signWordI := signWord(strs[i])
		if _, ok := hashMap[signWordI]; !ok {
			hashMap[signWordI] = []string{}
		}
		hashMap[signWordI] = append(hashMap[signWordI], strs[i])
	}

	for _, x := range hashMap {
		res = append(res, x)
	}

	return res
}

type SignWordType [26]int

func signWord(s string) SignWordType {
	m := [26]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := range s {
		m[s[i]-'a']++
	}

	return SignWordType(m)
}
