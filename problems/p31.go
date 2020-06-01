package problems

import "fmt"

// 降序排列最大
func nextPermutation(nums []int)  {
	nlen := len(nums)
	flag := false
	j := nlen - 1
	for ; j > 0 ; j-- {
		if nums[j] > nums[j-1] {
			flag = true
			break
		}
	}

	if !flag {
		// 数组已经是降序排列，最大的了，所以需要翻转到最小的
		i := 0
		j := nlen - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		return
	}

	pivot := j-1
	// pivot后的本来是倒序的，先把pivot后的翻转
	i := pivot + 1
	j = nlen - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	j = pivot + 1
	for j < nlen {
		if nums[j] > nums[pivot] {
			nums[pivot], nums[j] = nums[j], nums[pivot]
			break
		}
		j++
	}
}

func P31()  {
	a := []int{1,2,4,3,6,5}
	nextPermutation(a)
	fmt.Println(a)
	a = []int{6,5,4,3,2,1}
	nextPermutation(a)
	fmt.Println(a)
	a = []int{1,2,3,4,5,6}
	nextPermutation(a)
	fmt.Println(a)
}