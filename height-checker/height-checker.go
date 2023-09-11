func heightChecker(heights []int) int {
    if len(heights) < 2 {
        return 0
    }

    k := 0
    expected := make([]int, len(heights))

    for i:=0; i<len(heights); i++ {
        expected[i] = heights[i]
    }

    sorter(expected)

    for i:=0; i<len(heights); i++ {
        if heights[i] != expected[i] {
            k++
        }
    }

    return k
}

func sorter(nums []int) {
    for i:=0; i<len(nums); i++ {
        min := i
        for j:=i+1; j<len(nums); j++ {
            if nums[min] > nums[j] {
                min = j
            }
        }
        if i!=min {
            nums[i], nums[min] = nums[min], nums[i]
        }
    }
}
