package main

import (
	"fmt"
	"strconv"
	//"os"
)

func main() {
	var nombre string
	nombre = "Julian"
	fmt.Println(nombre)

	var x int
	x = 2
	fmt.Println(x)
	fmt.Printf("El valor de 'x' es %d\n", x)
	fmt.Printf("El valor de 'x' es %d y el nombre es %s\n", x, nombre)

	var bandera bool
	fmt.Println(bandera)

	var arr [3]int
	fmt.Println(arr)

	var decimal float32 = 11.141516789
	fmt.Printf("El valor de decimal es %.2f", decimal)

	y := 3
	fmt.Println("\nEl valor de 'y' es ", y)

	x = 5
	var strX string
	strX = strconv.Itoa(x)
	fmt.Println("El valor como texto de 'x' es " + strX)

	num := 1 // Con : se hace la declaracion y con = asignación de valor; el compilador asigna automat. el tipo.
	apellido := "Garcia"
	num = 2 // No es necesario volver a declarar con :, solo se aplica = para asignar nuevo valor.
	fmt.Printf("Num: es %d con Apellido: %s\n", num, apellido)

	const cte = 100
	//cte = 123 // No se puede cambiar el valor a una Constante.

}
