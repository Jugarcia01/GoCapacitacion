package main

import (
	"fmt"
)

func main() {
	//saludar("Julian")
	cadena := retornoConTipoStr()
	fmt.Println(cadena)

	numero := retornoConTipoInt()
	fmt.Println(numero)

	/*
		num, txt, flag := retornoCompuesto()
		fmt.Println(num)
		fmt.Println(txt)
		fmt.Println(flag)
	*/

	// Cuando solo necesito algun dato del retorno de la funcion y no se requieren los demás.
	num, _, _ := retornoCompuesto()
	fmt.Println(num)

}

func saludar(nombre string) {
	fmt.Println("Hola " + nombre)
}

func retornoConTipoStr() string {
	return "Hola mundo"
}

func retornoConTipoInt() int {
	return 123
}

func retornoCompuesto() (int, string, bool) {
	return 777, "Otro valor", true
}
