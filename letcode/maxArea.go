package letcode

func MaxArea(height []int) int {
	letf := 0
	right := len(height) - 1
	res := 0

	for letf < right {
		area := min(height[letf], height[right]) * (right - letf)

		if area > res {
			res = area
		}
		if height[letf] > height[right] {
			right--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
