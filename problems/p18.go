package problems

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	nlen := len(nums)
	if len(nums) < 4 {
		return ret
	}

	sort.Ints(nums)
	for i := 0; i < nlen - 3; i++ {
		for j := i+1; j < nlen - 2; j++ {
			ijsum := nums[i] + nums[j]
			p := j + 1
			q := nlen - 1
			for p < q {

				t := ijsum + nums[p] + nums[q]
				if t == target {
					ret = append(ret, []int{nums[i], nums[j], nums[p], nums[q]})
					// 去除相同的值
					for p++; p < q && nums[p] == nums[p-1]; {
						p++
					}
					for q--; p < q && nums[q] == nums[q+1]; {
						q--
					}
				} else if t > target {
					q--
				} else {
					p++
				}
			}
		}
	}

	return ret
}

func FourSum()  {
	// -2 -1 0 0 1 2
	a := []int{1, 0, -1, 0, -2, 2}
	b := fourSum(a, 3)
	fmt.Println(a)
	fmt.Println(b)
}