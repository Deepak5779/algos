package quickSortBased

func findKthLargestNumberInArray(nums []int, k int) int {
	return findKthLargestNumberInArrayQuickSortWay(nums, k, 0, len(nums)-1)
}

func findKthLargestNumberInArrayQuickSortWay(nums []int, k int, left int, right int) int {
	pivotIndex := getPivotIndex(nums, left, right)
	if pivotIndex == len(nums)-k {
		return nums[pivotIndex]
	} else if pivotIndex > len(nums)-k {
		return findKthLargestNumberInArrayQuickSortWay(nums, k, left, pivotIndex-1)
	} else {
		return findKthLargestNumberInArrayQuickSortWay(nums, k, pivotIndex+1, right)
	}
}

func getPivotIndex(nums []int, left int, right int) int {
	pIndex := left
	pValue := nums[pIndex]
	left++

	for left < right {
		for left < right && nums[left] <= pValue {
			left++
		}

		for left < right && nums[right] > pValue {
			right--
		}

		if left < right {
			swap(nums, left, right)
			left++
			right--
		}
	}

	// if left >= len(nums) {
	// 	left--
	// }

	// if right < 0 {
	// 	right++
	// }

	if left >= right {
		left--
	}

	swap(nums, pIndex, left)
	pIndex = left
	return pIndex
}

func swap(nums []int, i int, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
