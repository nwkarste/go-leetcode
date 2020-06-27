package utils

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	leftMost := nums[0]
	for left <= right {
		pos := (right-left)/2 + left
		val := nums[pos]
		if val == target {
			return pos
		}
		if val < target {
			if val < leftMost {
				if target < leftMost {
					left = pos + 1
				} else {
					right = pos - 1
				}
			} else {
				left = pos + 1
			}
		} else {
			if val < leftMost {
				right = pos - 1
			} else {
				if target < leftMost {
					left = pos + 1
				} else {
					right = pos - 1
				}
			}
		}
	}
	return -1
}
