func removeElement(nums []int, val int) int {
    var i, j, k = 0, len(nums)-1, 0

    for i<=j {
        if nums[j] == val {
            j--
            k++
            continue
        }
        if nums[i] == val {
            nums[i], nums[j] = nums[j], nums[i]
            j--
            k++
        }
        i++
    }

    return len(nums) - k
}