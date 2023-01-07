package letcode

func ContainsDuplicate(nums []int) bool {
	num_map := map[int]int{}

	for _, n := range num_map {
		if _, ok := num_map[n]; !ok {
			num_map[n] = 1
		} else {
			return true
		}
	}
	return false
}
