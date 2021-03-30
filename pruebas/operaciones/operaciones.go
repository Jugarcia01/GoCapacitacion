package main

import (
	"errors"
)

var err error

func main() {

}

func sumar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No pueden ser menores a cero")
	} else {
		resultado = x + y
	}
	return
}

func restar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No pueden ser menores a cero")
	} else if y > x {
		err = errors.New("Y no puede ser mayor que X")
	} else {
		resultado = x - y
	}
	return
}

func multip(x, y int) (resultado int, err error) {

	if x < 0 || y < 0 {
		err = errors.New("No pueden ser menores a cero")
	} else {
		resultado = x * y
	}
	return

}

func dividir(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No pueden ser menores a cero")
	} else if y == 0 {
		err = errors.New("ERROR: División por Cero dá un resultado indeterminado")
	} else {
		resultado = x / y
	}
	return
}
