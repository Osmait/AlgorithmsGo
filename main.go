package main

import (
	"fmt"

	"github.com/osmait/algorithmigo/challenger"
	linkedlist "github.com/osmait/algorithmigo/linkedList"
)

func main() {

	list := linkedlist.NewSingly[int]()

	err := list.AddAll("front", 1, 2, 3, 4, 5)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(challenger.IterativeFactorial(5))
	fmt.Println(challenger.RecurFactorial(5))
	challenger.RecurPermutation("abc", "")
	total := 1 << 5 //1 * 2^5
	fmt.Println(total)

	// list.AddAll("fdf", 1, 2, 3, 4, 5)
	// head := list.Head.Val
	// nextValue := list.Head.Next.Val
	// fmt.Println(head + nextValue)
	list.Display()
}
