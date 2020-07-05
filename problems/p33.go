package problems

import "fmt"

//https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			if nums[mid] > nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func P33()  {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
}
