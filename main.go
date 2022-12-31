package main

import (
	"fmt"

	"github.com/osmait/algorithmigo/sorting"
)

func main() {

	lista := []int{5, 4, 7, 10, 6}
	listaOrdenada := sorting.Bubble(lista)
	fmt.Println(listaOrdenada)
}
