// 1. Dado el radio de un círculo, calcular su perímetro y su área.

package main

import (
	"errors"
	"fmt"
	"strconv"
)

const miPi = 3.141592

func main() {
	var valor string
	var rad float64

	fmt.Println("CALCULO DE PERIMETRO Y AREA DE UN CIRCULO DADO EL VALOR DEL RADIO")
	fmt.Println("Digite el valor del Radio: ")

	fmt.Scanf("%s\n", &valor)
	rad, _ = strToFloat(valor)

	perimetro(rad)
	area(rad)
}

func strToFloat(dato string) (val float64, err error) {

	if val, err := strconv.ParseFloat(dato, 32); err == nil {
		fmt.Printf("%T, %v\n", dato, val)
	}
	return
}

func perimetro(rad float64) (perim float64, err error) {
	perim = 0.00

	if rad <= 0 {
		err = errors.New("El valor ingresado No puede ser inferior a 0")
	} else {
		perim = 2.00 * miPi * rad
		fmt.Printf("El perimetro del circulo es: %0.2f\n", perim)
	}
	return
}

func area(rad float64) (area float64, err error) {
	area = 0.00

	if rad <= 0 {
		err = errors.New("El dato ingresado No puede ser inferior a 0")
	} else {
		area = (miPi * rad * rad)
		fmt.Printf("El área del circulo es: %0.2f\n", area)
	}
	return
}
