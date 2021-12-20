package problems

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

func nSumTarget(nums []int, n, start, target int) [][]int {
	var res [][]int
	sizen := len(nums)
	if sizen < n || n < 2 {
		return res
	}
	if n == 2 {
		left := start
		right := sizen - 1
		for left < right {
			leftv := nums[left]
			rightv := nums[right]
			sum := leftv + rightv
			if sum == target {
				res = append(res, []int{nums[left], nums[right]})
				for left < right && nums[left] == leftv {
					left++
				}
				for left < right && nums[right] == rightv {
					right--
				}
			} else if sum < target {
				for left < right && nums[left] == leftv {
					left++
				}
			} else {
				for left < right && nums[right] == rightv {
					right--
				}
			}
		}

		return res
	}

	// n > 2
	for i := start; i < sizen; i++ {
		tmp := nSumTarget(nums, n-1, i+1, target-nums[i])
		for _, v := range tmp {
			v := append(v, nums[i])
			res = append(res, v)
		}
		for i < sizen-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}
