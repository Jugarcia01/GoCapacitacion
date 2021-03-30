// 8. Saber si un número es palíndromo.

package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int

	fmt.Print("Ingrese un numero entero positivo: ")
	fmt.Scan(&num)

	esPalindromo(num)

}

func esPalindromo(num int) (res bool, err error) {
	var tmpNumber, remainder int
	var reverse int = 0
	tmpNumber = num

	if num >= 0 {

		// Ciclo empleado en forma de While Loop
		for {
			remainder = tmpNumber % 10
			reverse = reverse*10 + remainder
			//fmt.Println(reverse)
			tmpNumber /= 10

			if tmpNumber == 0 {
				break // Rompe la ejecución del ciclo
			}
		}

		if num == reverse {
			fmt.Println(reverse)
			fmt.Printf("%d es un numero Palindromo", num)
			res = true
		} else {
			fmt.Println(reverse)
			fmt.Printf("%d No es un Palindromo", num)
			res = false
		}

	} else {
		err = errors.New("El valor ingresado no puede ser menor a Cero | 0")
	}
	return
}
