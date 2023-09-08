func sortArrayByParity(nums []int) []int {
    var sortedId, i int = 0, len(nums)-1

    for i>=0 && i>sortedId {
        if nums[sortedId]%2 == 0 {
            sortedId++
            continue
        }

        if nums[i]%2 == 0 {
            nums[sortedId], nums[i] = nums[i], nums[sortedId]
            sortedId++
            continue
        }

        i--
    }

    return nums
}