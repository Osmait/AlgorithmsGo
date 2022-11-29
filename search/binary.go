package search

func Binary(array []int, target, lowIndex, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, ErrNotFound
	}
	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

func BinaryIterrative(array []int, target int) (int, error) {
	starIndex := 0
	endIndex := len(array) - 1
	var mid int
	for starIndex <= endIndex {
		mid = int(starIndex + (endIndex-starIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			starIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}
