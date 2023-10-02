func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))

    mid := middle(nums, 0, len(nums)-1, minIdx)

    var way int
    for k, l, r := 0, mid, mid+1; (l>=0 || r<len(nums)) && k<len(res); k++ {
        left, right := -1, -1

        if l>=0 {
            left = sqrt(nums[l])
        }

        if r<len(nums) {
            right = sqrt(nums[r])
        }

        if left != -1 && right != -1 {
            res[k], way = minX(left, right)
            if way > 0 {
                r++
            }else{
                l--
            }
            continue
        }

        if left == -1 && right != -1 {
            res[k] = right
            r++
            continue
        }

        if left != -1 && right == -1 {
            res[k] = left
            l--
            continue
        }

        break
    }

    return res
}

func middle(nums []int, left, right int, f func([]int, int, int) int) int {
    if right-left <= 1 {
        return f(nums, left, right)
    }

    mid := (left+right)/2
    if left >= 0 {
        left = middle(nums, left, mid, f)
    }

    if right < len(nums) {
        right = middle(nums, mid, right, f)
    }

    return f(nums, left, right)
}

func minIdx(nums[]int, l, r int) int {
    if abc(nums[l]) < abc(nums[r]) {return l}
    return r
}

func minX(a, b int) (int, int) {
    if a<b {return a, -1}
    return b, 1
}

func abc(num int) int {
    if num > 0 {return num}
    return -num
}

func sqrt(num int) int {
    return num*num
}