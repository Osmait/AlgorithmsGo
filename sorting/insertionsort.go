package sorting

import (
	"fmt"

	"github.com/osmait/algorithmigo/constraints"
)

func Insertion[T constraints.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex
		for ; iterator > 0 && arr[iterator-1] > temporary; iterator-- {
			fmt.Println(arr[iterator])
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}
