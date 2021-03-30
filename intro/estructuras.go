package main

import (
	"fmt"
	"strconv"
)

//Se crea una estructura con sus propiedades y sus respectivos tipos.
type Humano struct {
	nombre string
	edad   int
}

type Grupo struct {
	integrantes []Humano
}

func (humano Humano) imprimir() string {
	return "\nNombre: " + humano.nombre + ", edad: " + strconv.Itoa(humano.edad)
}

func (h *Humano) cambiarValoresPunt(nombre string, edad int) { // Aquí se pasa la posic. de memoria de la estructura
	h.nombre = nombre
	h.edad = edad
}

func (h Humano) cambiarValores(nombre string, edad int) { // Aquí se pasa es una copia/valor de la estructura, pero no la direccion.
	h.nombre = nombre
	h.edad = edad
}

func main() {
	persona := Humano{"Julian", 38}
	fmt.Println(persona)
	fmt.Println(persona.imprimir())

	persona2 := new(Humano)
	fmt.Println(persona2)

	persona2 = new(Humano)
	fmt.Println(persona2.imprimir())
	persona2.nombre = "Pedro"
	persona2.edad = 20

	persona.nombre = "David"
	persona.edad = 10
	fmt.Println(persona.imprimir())
	fmt.Println(persona.nombre)

	fmt.Printf("%+v\n", persona)

	persona.cambiarValores("Pablo", 32)
	fmt.Printf(persona.imprimir())

	persona2.cambiarValores("Jose", 37)
	fmt.Printf(persona2.imprimir())

	persona2.cambiarValoresPunt("Jose", 37)
	fmt.Printf(persona2.imprimir())

}
