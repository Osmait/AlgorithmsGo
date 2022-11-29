package main

import (
	"fmt"

	"github.com/osmait/algorithmigo/search"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, _ := search.Binary(arr, 3, 0, len(arr))
	fmt.Println(result)
}
