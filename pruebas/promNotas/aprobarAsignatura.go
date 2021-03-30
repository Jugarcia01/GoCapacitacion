// 5. Calcular la nota necesaria para aprobar una asignatura, dadas las tres primeras notas.
// Las 4 notas tienen el mismo porcentaje.

package main

import (
	"errors"
	"fmt"
)

func main() {
	var notaMinApr = 3.0
	var notas []float64 = []float64{2.5, 2.5, 2.5}

	nota, err := calcNotaAprob(notas, notaMinApr)
	if err == nil {
		err = errors.New("El valor ingresado No puede menor que Cero | 0")
	}
	fmt.Println("Nota min requerida:", nota)
}

func calcNotaAprob(notas []float64, notaMinApr float64) (notaMinAprobac float64, err error) {
	var sumNotas float64
	var cantNotas float64
	var dataErr = 0
	var res float64

	if len(notas) == 0 {
		dataErr++
		err = errors.New("La cantidad de notas no puede ser Cero | 0")

	} else {
		for _, v := range notas {
			if v < 0.0 {
				err = errors.New("Las notas No pueden ser menores a Cero | 0")
				dataErr++
			}
		}
	}

	if dataErr == 0 {
		for _, v := range notas {
			sumNotas += v
		}
		cantNotas = float64(len(notas) + 1)

		res = (notaMinApr * cantNotas) - sumNotas

		if res <= 0.0 {
			notaMinAprobac = 0.0
		} else if res >= 5.0 {
			notaMinAprobac = res
			err = errors.New("El estudiante reprobó la materia :-(")
		} else {
			notaMinAprobac = res
		}
	}
	return
}
