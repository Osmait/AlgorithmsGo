package main

import (
	"fmt"

	linkedlist "github.com/osmait/algorithmigo/linkedList"
)

func main() {

	var i int

	for {
		fmt.Println("1-Ingresar un Dato: ")
		fmt.Println("2-Eliminar al primero ")
		fmt.Println("3-Ver  Almacen ")

		fmt.Scan(&i)
		almancen := linkedlist.NewSingly[int]()
		almancen.Display()

		switch i {
		case 1:

		case 2:
			almancen.DelAtBeg()
		case 3:
			almancen.Display()

		default:
			fmt.Println("opcion invalida")

		}
	}

}
