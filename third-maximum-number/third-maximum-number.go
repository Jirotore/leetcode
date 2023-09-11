func thirdMax(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    if len(nums) == 2 {
        if nums[0] < nums[1] {return nums[1]}
        return nums[0]
    }

    sorter(nums)

    k:=3
    max:=nums[0]
    for i:=0; i<len(nums) && k>0; i++ {
        if i+1<len(nums) && nums[i] == nums[i+1] {
            continue
        }
        k--
        max=nums[i]
    }
    
    if k > 0 {
        max = nums[0]
    }

    return max
}

func sorter(nums[]int) {
    for i:=0; i<len(nums);i++ {
        min:=i
        for j:=i+1; j<len(nums); j++ {
            if nums[min]<nums[j] {
                min=j
            }
        }
        if i!=min {
            nums[i], nums[min] = nums[min], nums[i]
        }
    }
}
