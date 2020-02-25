package problem0034

func searchRange(nums []int, target int) []int {
	l := First(nums, target)
	if l == -1 {
		return []int{-1, -1}
	}
	last := Last(nums, target, l)
	return []int{l, last}
}

func First(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 1)
		switch {
		case nums[mid] < target:
			l = mid + 1
		case nums[mid] > target:
			r = mid - 1
		default:
			if (mid == 0) || (nums[mid-1] != target) {
				return mid
			}
			r = mid - 1
		}
	}
	return -1
}

func Last(nums []int, target int, l int) int {
	r := len(nums) - 1
	for l <= r {
		mid := l + ((r - l) / 2)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
