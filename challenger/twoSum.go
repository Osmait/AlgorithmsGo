package challenger

func TwoSum(nums []int, target int) []int {
	var lista []int
	for i, v := range nums {
		for j, n := range nums {
			if v+n == target && i != j {
				lista = []int{i, j}
			}
		}

	}
	return lista
}
