package problems

import "fmt"
// https://leetcode-cn.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	nlen := len(nums)
	total := 1
	index0 := []int{}
    for i := 0; i < nlen; i++ {
		if nums[i] == 0 {
			index0 = append(index0, i)
		} else {
			total *= nums[i]
		}
	}

	result := []int{}
	len0 := len(index0)
	for i := 0; i < nlen; i++ {
		if len0 > 0 {
			if len0 == 1 && nums[i] == 0 {
				result = append(result, total)
			} else {
				result = append(result, 0)
			}
		} else {
			result = append(result, total/nums[i])
		}
	}

	return result
}

func P238()  {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
	fmt.Println(productExceptSelf([]int{0,0}))
}
