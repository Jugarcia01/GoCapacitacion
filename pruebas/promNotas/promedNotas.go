// 2. Calcular el promedio de N notas.

package main

import (
	"errors"
)

func main() {

	var notas []float64

	// notas = []float64{}
	// fmt.Println(len(notas))

	promedNotas(notas)

}

func promedNotas(notas []float64) (prom float64, err error) {
	var num float64 = 0.0
	var valInvalid int = 0

	if notas == nil {
		err = errors.New("No existen notas para promediar")
	} else if len(notas) <= 0 {
		err = errors.New("No existen notas a las cuales calcular el promedio")
	} else {
		for _, n := range notas {
			if n >= 0 {
				num = num + n
			} else {
				valInvalid++
				err = errors.New("Las notas no pueden ser inferiores a 0")
			}
		}

		if valInvalid == 0 && num >= 0.00 {
			prom = num / float64(len(notas))
		}
	}
	return
}
