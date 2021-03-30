package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string

	//scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("Ingrese una frase: ")
	//scanner.Scan()
	//frase = scanner.Text()
	frase = "in on of in on on, of on in on"
	fmt.Printf("Frase ingresada: %s\n", frase)

	var wordsMap map[int]string
	wordsMap = separaPalabras(frase)
	fmt.Println(wordsMap)
}

func separaPalabras(txt string) map[int]string {
	mapFrase := make(map[int]string)
	palabras := strings.Fields(txt)

	for ind, palabra := range palabras {
		fmt.Printf("Palabra %d es: %s\n", ind, palabra)
		mapFrase[ind] = palabra
	}
	return mapFrase
}
