package sorting

import (
	"fmt"

	"github.com/osmait/algorithmigo/constraints"
)

func Bucket[T constraints.Number](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	// Find the maximum and minimun elements in arr
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	// Create an empty bucket for each element in arr
	bucket := make([][]T, len(arr))

	// put each element in the appropiate bucket
	for _, v := range arr {
		bucketIndex := int((v - min) / (max - min) * T(len(arr)-1))
		fmt.Println(bucketIndex)
		bucket[bucketIndex] = append(bucket[bucketIndex], v)
		fmt.Println(bucket[bucketIndex])
	}
	for i := range bucket {
		bucket[i] = Insertion(bucket[i])
	}
	sorted := make([]T, 0, len(arr))
	for _, v := range bucket {
		sorted = append(sorted, v...)

	}
	return sorted

}
