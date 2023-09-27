func findDisappearedNumbers(nums []int) []int {
    var lost int

    for i:=0; i<len(nums); {
        if nums[i] != i+1 && nums[i] != 0 {
            if nums[nums[i]-1] != nums[i] {
                nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
                continue
            }
            nums[i] = 0
            lost++
        }

        i++
    }

    for i, k:=0, 0; i<len(nums) && k<lost; i++ {
        if nums[i] == 0 {
            nums[k] = i+1
            k++
        }
    }

    return nums[:lost]
}