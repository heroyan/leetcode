package problems

import (
	"fmt"
	"reflect"
	"strconv"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 归并排序后再求中位数
	l1 := len(nums1)
	l2 := len(nums2)
	var tmp []int = []int{}
	i := 0
	j := 0
	for i < l1 && j < l2 {
		if nums1[i] < nums2[j] {
			tmp = append(tmp, nums1[i])
			i++
		} else {
			tmp = append(tmp, nums2[j])
			j++
		}
	}

	for i < l1 {
		tmp = append(tmp, nums1[i])
		i++
	}

	for j < l2 {
		tmp = append(tmp, nums2[j])
		j++
	}

	tlen := len(tmp)
	result := 0.0
	if tlen % 2 == 1 {
		result = float64(tmp[tlen/2])
	} else {
		result = float64((tmp[tlen/2] + tmp[tlen/2-1]))/2
	}

	v, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", result), 64)

	return v
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n % 2 == 1 {
		return float64(getKthElem(nums1, nums2, n/2+1))
	}

	return float64(getKthElem(nums1, nums2, n/2+1) + getKthElem(nums1, nums2, n/2-1+1)) / 2
}

func getKthElem(nums1 []int, nums2 []int, k int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	start1 := 0
	start2 := 0
	for  {
		if start1 == len1 {
			return nums2[start2+k-1]
		}
		if start2 == len2 {
			return nums1[start1+k-1]
		}
		if k == 1 {
			return min(nums1[start1], nums2[start2])
		}
		half := k/2
		newstart1 := min(start1+half, len1) - 1
		newstart2 := min(start2+half, len2) - 1
		privot1, privot2 := nums1[newstart1], nums2[newstart2]
		if privot1 <= privot2 {
			k -= (newstart1-start1+1)
			start1 = newstart1 + 1
		} else {
			k -= (newstart2-start2+1)
			start2 = newstart2 + 2
		}
	}

	return 0
}

func min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func TestArray(arr [10]int)  {
	arr[0] = 100
}

func TestSlice(arr []int)  {
	arr[0] = 100
}

type ttt struct {
	a string
}

func (t *ttt) test() {
	fmt.Println(t.a)
}


func TestA()  {
	var b [10]int = [10]int{1,2,3,4,5}
	TestArray(b)
	fmt.Println(b)
	a := []int{1,2,3,4,5}
	TestSlice(a)
	fmt.Println(a[5:])
	fmt.Println(cap(a))
	var t interface{} = ttt{a:"hi"}
	fmt.Println(reflect.TypeOf(t))
}

func quikSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quikSort(left)
	right = quikSort(right)

	return append(append(left, pivot), right...)
}

func quikSort1(arr []int, left, right int) {
	pivot := arr[left]
	pindex := left
	i := left
	j := right
	for i <= j {
		for pindex <= j && arr[j] >= pivot  {
			j--
		}
		if pindex <= j {
			arr[pindex] = arr[j]
			pindex = j
		}
		for i <= pindex && arr[i] <= pivot {
			i++
		}
		if i <= pindex {
			arr[pindex] = arr[i]
			pindex = i
		}
	}
	arr[pindex] = pivot

	if pindex > left + 1 {
		quikSort1(arr, left, pindex - 1)
	}

	if pindex + 1 < right {
		quikSort1(arr, pindex + 1, right)
	}

	return
}

func TestQ()  {
	arr := []int{12,123,123,4,1,3,66,2,1,4,1,55,2344,5}
	quikSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)
}