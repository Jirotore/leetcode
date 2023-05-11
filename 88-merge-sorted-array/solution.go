package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	var buf = make([]int, m+n)

	var k int
	for i, j := 0, 0; i < m || j < n; k++ {
		if (j >= n) || (i < m && nums1[i] < nums2[j]) {
			buf[k] = nums1[i]
			i++
		} else {
			buf[k] = nums2[j]
			j++
		}
	}

	for i := 0; i < m+n; i++ {
		nums1[i] = buf[i]
	}
}
