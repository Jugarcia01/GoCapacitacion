package main

import (
	"testing"
)

func TestAprobarAsignatura(t *testing.T) {

	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := calcNotaAprob([]float64{5, 4.5, 4.0}, 3.0)
		esperado := (3.0 * 4.0) - (5.0 + 4.5 + 4.0)

		if resultado != esperado { //math.Trunc(esperado)
			t.Errorf("Resultado : %f, se esperaba: %f", resultado, esperado)
		}
	})

	t.Run("Probar cuando solo hay una nota y es superior a 0", func(t *testing.T) {
		resultado, _ := calcNotaAprob([]float64{3.0}, 3.0)
		esperado := (3.0 * 2.0) - (3.0)

		if resultado != esperado {
			t.Errorf("Resultado : %f, se esperaba: %f", resultado, esperado)
		}
	})

	t.Run("Probar que Existe error cuando no hay notas", func(t *testing.T) {
		_, err := calcNotaAprob([]float64{}, 3.0)

		if err == nil {
			t.Errorf("Debe haber almenos una nota para determnar la nota de la asignatura")
		}
	})

	t.Run("Probar que las notas sean superiores a 0", func(t *testing.T) {
		_, err := calcNotaAprob([]float64{-5.00, 0.00, -1.22, 2.11}, 3.0)

		if err == nil {
			t.Errorf("Las notas deben ser mayores o iguales a 0.00")
		}
	})

}
