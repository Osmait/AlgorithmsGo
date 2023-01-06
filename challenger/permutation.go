package challenger

import (
	"fmt"
)

func RecurPermutation(str, pocket string) {
	if pocket == "" {
		pocket = ""
	}
	if len(str) == 0 {
		fmt.Println(pocket)
	} else {
		for i := 0; i < len(str); i++ {
			letter := string(str[i])
			front := str[0:i]
			back := str[i+1:]
			together := front + back
			RecurPermutation(together, letter+pocket)
		}
	}

}
