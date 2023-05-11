package leetcode

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int, k)
	for _, num := range nums {
		freqMap[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for k, v := range freqMap {
		buckets[v] = append(buckets[v], k)
	}

	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(res) == k {
			break
		}
		res = append(res, buckets[i]...)
	}

	return res
}
