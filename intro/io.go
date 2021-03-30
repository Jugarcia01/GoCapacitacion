package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nombre string
	var edad int
	saludar(nombre)

	fmt.Println("Ingresa tu nombre: ")
	//fmt.Scanf("%s", &nombre) // OJO: NO TOMA CADENAS CON ESPACIO, Cuando se ingresa un espacio lo toma como un nuevo argumento.
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	line := in.Text()
	fmt.Println("Su nombre es: ", line)

	fmt.Println("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad) // OJO: Cuando hay dos Scanf PASA POR ALTO EL ULTIMO (PQ PASA DERECHO?).
	fmt.Printf("Tu edad es: %d años.", edad)
}

func saludar(nombre string) {
	fmt.Println("Hola bienvenido!" + nombre)
}
