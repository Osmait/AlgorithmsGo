package main

import (
	"fmt"

	"github.com/osmait/algorithmigo/challenger"
	linkedlist "github.com/osmait/algorithmigo/linkedList"
)

func main() {

	// arr := []int{2, 5, 5, 11}
	// result, _ := search.Binary(arr, 3, 0, len(arr))
	// fmt.Println(result)
	// result := challenger.TwoSum(arr, 10)
	// fmt.Println(result)
	fmt.Println(challenger.Palindromo(121))
	prueba := linkedlist.NewSingly[int]()
	prueba.AddAtBeg(1)
	prueba.AddAtBeg(10)
	prueba.AddAtBeg(13)
	prueba.AddAtBeg(177)

	prueba.Display()

}
