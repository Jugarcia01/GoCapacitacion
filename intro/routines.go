package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Imprimiendo nombre ...")
	go imprimirNombre("Julian") // gorutina que se ejecuta en paralelo en otro contexto "hilo" de ejecucion
	fmt.Println("Fin de impresion")
	time.Sleep(1 * time.Second)
	//var espera string
	//fmt.Scanln(&espera) // Se plantea espera para poder dar tiempo a observar lo que sucede con la gorutina.
}

func imprimirNombre(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(letra)
	}
}
