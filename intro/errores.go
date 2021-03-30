package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := operar(-1, 2, "resta")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

}

//func operar(x, y int, operac, otraCadena string) int {
func operar(x, y int, operac string) (int, error) {
	var err error

	// Validando que los numeros no sean menores a Cero
	if x < 0 || y < 0 {
		err = errors.New("La variable es menor a 0")
		return 0, err
	}

	switch operac {
	case "suma":
		return x + y, nil
	case "resta":
		return x - y, nil
	case "multiplicacion":
		return x * y, nil
	case "division":
		return x / y, nil
	default:
		return x, nil
	}
}
