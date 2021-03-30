package main

import (
	"fmt"
	"reflect"
)

type minuto int // Se crea un nuevo tipo de dato llamado minuto
func (m minuto) aInt() int {
	return int(m) * 10
}

type hora int

func main() {

	var tiempo minuto
	fmt.Println(tiempo)
	fmt.Println(reflect.TypeOf(tiempo)) // Para imprimir/observar el tipo de una variable se emplea fmt.Println(reflect.TypeOf(VARIABLE))

	num := 3
	tiempo = minuto(num)
	fmt.Println(num, tiempo)

	var time minuto = 7
	fmt.Println(time)
	var convertido int
	convertido = time.aInt()
	fmt.Println(convertido)

	var horas hora
	fmt.Println(horas)
	// fmt.Println(horas.aInt()) // ERROR: hora.aInt undefined (type hora has no method aInt)

}
