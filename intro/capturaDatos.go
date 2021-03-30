package main

import (
	"fmt"
	"strconv"
)

const cantEst = 3

func main() {
	var nombre [cantEst]string
	var edad [cantEst]int
	var nota [cantEst]int
	var edadProm float64
	var notaProm float64
	var isNum bool

	fmt.Println("SISTEMA DE INGRESO DE DATOS DE ESTUDIANTES")
	for i := 0; i < len(edad); i++ {

		fmt.Println("Ingrese nombre del estudiante: ")
		fmt.Scanf("%s\n", &nombre[i])
		fmt.Printf("Nombre ingresado: %s\n", nombre[i])

		isNum = false
		for !isNum {
			fmt.Println("Ingrese edad del estudiante: ")
			var edadS string
			fmt.Scanf("%s\n", &edadS)
			datoIn, err := strconv.Atoi(edadS)
			if err != nil {
				fmt.Println("ERROR: Debe ingresar un número de Edad válido.")
				isNum = false
			} else {
				edad[i] = datoIn
				fmt.Printf("Edad: %d años.\n", edad[i])
				isNum = true
			}
		}

		isNum = false
		for !isNum {
			fmt.Println("Ingrese nota del estudiante: ")
			var notaS string
			fmt.Scanf("%s\n", &notaS)
			notaIn, err := strconv.Atoi(notaS)
			if err != nil {
				fmt.Println("ERROR: Debe ingresar un número de Nota válido.")
				isNum = false
			} else {
				nota[i] = notaIn
				fmt.Printf("Nota: %d puntos.\n", nota[i])
				isNum = true
			}
		}

	}

	edadProm = promedio(edad)
	fmt.Printf("\nEl promedio de edad de los estudiantes es: %.2f\n", edadProm)

	notaProm = promedio(nota)
	fmt.Printf("El promedio de notas de los estudiantes es: %.2f\n", notaProm)

}

func promedio(dato [cantEst]int) (result float64) {
	num := 0.0
	den := 1.0

	for i := 0; i < len(dato); i++ {
		num = num + float64(dato[i])
	}
	den = float64(len(dato))
	return (num / den)
}
