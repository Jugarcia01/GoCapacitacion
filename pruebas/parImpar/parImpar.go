// 4. Saber si un número es par o impar.

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var in string

	fmt.Printf("Digite un numero: ")
	fmt.Scanf("%s\n", &in)
	valor, err := strconv.Atoi(in)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		parImparAcotado(valor)
	}
}

func parImparAcotado(num int) (res bool, err error) {

	if num < 0 {
		err = errors.New("El valor ingresado No puede menor que Cero | 0")
	} else if num > 100000 {
		err = errors.New("El valor ingresado No puede ser superior a Cien Mil | 100.000")
	} else {
		if num%2 == 0 {
			res = true
		} else {
			res = false
		}
	}
	return
}

// func parImpar(num int) (res bool, err error) {
// 	if num%2 == 0 {
// 		// fmt.Println("El numero: ", num, " es Par")
// 		res = true
// 	} else {
// 		// fmt.Println("El numero: ", num, " es Impar")
// 		res = false
// 	}
// 	return
// }
