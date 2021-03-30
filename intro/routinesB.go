package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inicio")
	for i := 1; i <= 20; i++ {
		fmt.Println("Lanzando Go routine #", i)
		//time.Sleep(100 * time.Millisecond)
		go func(num int) {
			//time.Sleep(100 * time.Millisecond)
			fmt.Printf("Go routine # %d\n", num)
		}(i) // Este valor de i solo se puede modificar si se pasa como puntero, de lo contrario NO se puede modificar.
	}
	fmt.Println("Fin")
	var espera string
	fmt.Scanln(&espera)
}
