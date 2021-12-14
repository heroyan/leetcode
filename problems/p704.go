package problems

func search2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func searchRange2(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	leftArea := -1
	// 求左边界
	for left <= right {
		mid := left + (right-left) / 2
		if nums[mid] == target {
			right = mid - 1
			leftArea = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 求右边界
	left = 0
	right = len(nums) - 1
	rightArea := -1
	// 求右边界
	for left <= right {
		mid := left + (right-left) / 2
		if nums[mid] == target {
			left = mid + 1
			rightArea = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{leftArea, rightArea}
}

func searchRange3(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	leftArea := -1
	// 求左边界
	for left <= right {
		mid := left + (right-left) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left >= len(nums) || nums[left] != target {
		leftArea = -1
	} else {
		leftArea = left
	}

	// 求右边界
	left = 0
	right = len(nums) - 1
	rightArea := -1
	// 求右边界
	for left <= right {
		mid := left + (right-left) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if right < 0 || nums[right] != target {
		rightArea = -1
	} else {
		rightArea = right
	}

	return []int{leftArea, rightArea}
}