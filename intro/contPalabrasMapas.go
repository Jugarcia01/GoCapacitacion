package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var frase string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese una frase: ")
	scanner.Scan()
	frase = scanner.Text()
	//frase = "in on go in go on, of on in go go"
	//fmt.Printf("Frase ingresada: %s\n", frase)

	wordsCont := contPalabras(frase)
	fmt.Println(wordsCont)

	palMasRepet, cantidad := maxFreq(wordsCont)
	fmt.Printf("\nPalabra más repetida: %s\n", palMasRepet)
	fmt.Printf("Num Repeticiones: %d\n", cantidad)
}

func contPalabras(txt string) map[string]int {
	palabras := strings.Fields(txt)     // El texto es convertido a un array de strings, separa cada palabra en un item del vector palabras
	mapPalabras := make(map[string]int) // Almacenara las palabras y cantidad de veces que existe cada palabra => Mapa[palabra]cantVecesPalabra

	for _, palabra := range palabras { // Se recorre el vector de palabras y se va alimentando mapPalabras
		_, cant := mapPalabras[palabra]
		if cant { // Se compara si la palabra existe o no dentro de mapPalabras
			mapPalabras[palabra]++ // Si la palabra existe en mapPalabras entonces se va incrementando en 1 el valor
		} else {
			mapPalabras[palabra] = 1 // Si la palabra no existe en mapPalabras entonces se fija cantidad en 1
		}
	}
	return mapPalabras
}

func maxFreq(mapaPalab map[string]int) (string, int) {
	var palabMasRep string
	var numRep int

	for palabra, cant := range mapaPalab {
		if cant > numRep {
			numRep = cant
			palabMasRep = palabra
		}
	}
	return palabMasRep, numRep
}
