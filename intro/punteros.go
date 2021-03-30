package main

import "fmt"

func main() {
	num := 10
	fmt.Println(num)
	changeNum(num)
	fmt.Printf("Despues de la funcion queda el valor num: %d\n", num)

	var puntero *int      // Asi se crea un puntero de un tipo de dato
	puntero = &num        // Así se asigna la posic. de memoria de la variable num al puntero.
	fmt.Println(puntero)  // Imprime la direcc de la memoria donde esta ubicado el puntero
	fmt.Println(*puntero) // Imprime el valor o lo que hay almacenado en la posic. de memoria del puntero
	*puntero = 100        // Cambiando el valor del puntero.
	fmt.Println(*puntero)
	fmt.Printf("Despues del cambio del valor del puntero el valor de num: %d\n", num)

	txt := "Este texto se alamacena en una posicion de memoria."
	var punteroTxt *string
	punteroTxt = &txt        // Así se asigna la posic. de memoria de la variable num al puntero.
	fmt.Println(punteroTxt)  // Imprime la direcc de la memoria donde esta ubicado el puntero
	fmt.Println(*punteroTxt) // Imprime el valor o lo que hay almacenado en la posic. de memoria

	num2 := 10
	fmt.Printf("\nValor entregado: %d\n", num2)
	changeNumPunt(&num2) // Paso por referencia
	fmt.Printf("Despues de la funcion queda el valor num: %d\n", num2)

}

// Funcion normal paso por valor, se modifica el valor pero en una variable con un valor de memoria dif. al de la var original
func changeNum(num int) { // Aqui el valor de num del main se pasa a una copia de otra var num
	num = num * 1000 // Esta variable num es diferente a la declarada en el main
	fmt.Printf("En la func. changeNum existe el valor num: %d\n", num)

}

// Funcion con punteros, se modifica el valor que existe en la posicion de memoria.
func changeNumPunt(num *int) { // Aqui el valor de num del main se pasa la direccion de memoria de num
	*num = *num * 1000 // Aquí se esta modificando el valor que existe en la posic. de memoria que se pasí
	fmt.Printf("En la func. changeNum existe el valor num: %d\n", *num)
}
