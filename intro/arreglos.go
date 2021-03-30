package main

import "fmt"

func main() {
	var arreglo [5]int
	arreglo[0] = 12
	fmt.Println(arreglo)

	var arregloBool [5]bool
	fmt.Println(arregloBool)

	var arrInt [5]int
	for i := 0; i < len(arrInt); i++ {
		arrInt[i] = i * 10
	}
	fmt.Println(arrInt)

	var matriz [3][3]int
	for f := 0; f < len(matriz); f++ {
		for c := 0; c < len(matriz); c++ {
			matriz[f][c] = (f + c) * 10
		}
	}
	fmt.Println(matriz)

	arreglo3 := [3]string{"A", "B", "C"}
	arreglo3[1] = "Z"
	fmt.Println(arreglo3)

}
