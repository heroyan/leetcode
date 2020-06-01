package problems

import "fmt"

func searchRange(nums []int, target int) []int {
    left := 0
    right := len(nums) - 1
    start := -1
    end := -1
    for left <= right {
        mid := left + (right - left) / 2
        if nums[mid] == target {
            i := mid
            j := mid
            for ; i >= 0 && nums[i] == target; i-- {}
            for ; j <= len(nums)-1 && nums[j] == target; j++ {}
            start = i+1
            end = j-1
            break
        } else if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        }
    }

    return []int{start, end}
}

func P34()  {
    a := []int{5,7,7,8,8,10}
    b := searchRange(a, 8)
    fmt.Println(b)
}