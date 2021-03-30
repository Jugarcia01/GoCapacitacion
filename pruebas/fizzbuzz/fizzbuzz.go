// Aplicando TDD, se plantean los casos de pruebas, y de ahi se parte para ir codificando
// A medida que se va resolviendo cada caso de prueba, se va avanzando al siguiente
// hasta tener completo el desarrollo.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}

}

func fizzbuzz(numero int) string {
	if numero <= 0 {
		return "No es un núm válido"
	}

	if numero%3 == 0 && numero%5 == 0 {
		return "fizzbuzz"
	}

	if numero%3 == 0 {
		return "fizz"
	}

	if numero%5 == 0 {
		return "buzz"
	}

	//return string(rune(numero))
	return strconv.Itoa(numero)
}
