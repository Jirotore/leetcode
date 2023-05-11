package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	left, right, mid := 0, len(matrix)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if matrix[mid][0] <= target && target <= matrix[mid][len(matrix[mid])-1] {
			return binnarySearch(matrix[mid], target)
		}
		if matrix[mid][0] < target && target > matrix[mid][len(matrix[mid])-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func binnarySearch(arr []int, target int) bool {
	left, right, mid := 0, len(arr)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
