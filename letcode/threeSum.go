package letcode

import "sort"

func thrreSum(nums []int) [][]int {
	n := len(nums)

	sort.Ints(nums)

	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}
		num2Idx := num1Idx + 1
		num3Idx := n - 1

		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				result = append(result, []int{nums[num2Idx], nums[num3Idx]})
				num3Idx--

				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx] {
					num3Idx--
				}
			} else if sum > 0 {
				num3Idx--
			} else {
				num2Idx++
			}

		}
	}
	return result
}
