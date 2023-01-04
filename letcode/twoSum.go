package letcode

func TwoSum(nums []int, target int) [2]int {
	var ret [2]int
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == key {
				ret[0] = i
				ret[1] = j
			}
		}
	}
	return ret
}
