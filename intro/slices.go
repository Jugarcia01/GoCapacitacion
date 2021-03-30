// Son similares a las listas.
// Similar a los arreglos pero se puede cambiar su tamaño de forma dinamica.

package main

import "fmt"

func main() {
	var arreglo [5]int // Los Arrays tienen una longitud establecida.
	//var arreglo2 [10]int // Los Arrays tienen una longitud establecida.
	var slice []int // Los Slices tienen longitud variable.

	// Pasando los datos a un nuevo arreglo
	//copy(arreglo2, arreglo)
	//fmt.Println(arreglo2)

	fmt.Println(arreglo)
	fmt.Println(slice)

	if slice == nil {
		fmt.Println("El slice esta vacío.")
	} else {
		fmt.Println(slice)

	}

	slice = append(slice, 10)
	fmt.Println(slice)
	fmt.Println(slice[0])

	//slice[3] = 4  //panic: runtime error: index out of range [3] with length 1
	//fmt.Println(slice)

	fmt.Println(len(slice)) // Determina longitud del slice
	fmt.Println(cap(slice)) // Determina capacidad del slice

	slice2 := make([]int, 3, 10)
	slice2 = append(slice2, 7)

	fmt.Println("Verificando datos de los slice:")
	fmt.Println(slice)
	fmt.Println(slice2)

	fmt.Println("Verificando longitud de los slice:")
	fmt.Println(len(slice))
	fmt.Println(len(slice2))

	fmt.Println("Verificando capacidad de los slice:")
	fmt.Println(cap(slice))
	fmt.Println(cap(slice2))

	//sliceMat := make([][]int, 3, 10)
	//sliceMat = append(sliceMat, [0][0], 7)
	//fmt.Println(sliceMat)

	arregloTen := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// Para efectuar la copia de un arreglo en un slice
	//sliceTen := arregloTen[ini:(fin-1)]

	// Copiando un fragmento del arreglo en un slice
	sliceTen := arregloTen[3:6]
	fmt.Println(arregloTen)
	fmt.Println(sliceTen)

	// Copiando todo el arreglo en un slice
	sliceAll := arregloTen[:]
	fmt.Println(arregloTen)
	fmt.Println(sliceAll)

	fuente := []int{10, 20, 30, 40, 50, 60, 70}
	destino := make([]int, 5)
	fmt.Println(fuente)
	fmt.Println(destino)

	copy(destino, fuente) // Realizacion de copias entre slices
	fmt.Println(destino)

}
