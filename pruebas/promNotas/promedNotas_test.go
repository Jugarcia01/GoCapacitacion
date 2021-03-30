package main

import (
	"testing"
)

func TestPromedNotas(t *testing.T) {

	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := promedNotas([]float64{10.00, 20.00, 30.00})
		esperado := (10.00 + 20.00 + 30.00) / (3.00)

		if resultado != esperado { //math.Trunc(esperado)
			t.Errorf("Resultado : %f, se esperaba: %f", resultado, esperado)
		}
	})

	t.Run("Probar cuando solo hay una nota y es superior a 0", func(t *testing.T) {
		resultado, _ := promedNotas([]float64{77})
		esperado := (77.00 / 1.00)

		if resultado != esperado {
			t.Errorf("Resultado : %f, se esperaba: %f", resultado, esperado)
		}
	})

	t.Run("Probar que Existe error cuando no hay notas", func(t *testing.T) {
		_, err := promedNotas([]float64{})

		if err == nil {
			t.Errorf("Debe haber almenos una nota para promediar")
		}
	})

	t.Run("Probar que las notas sean superiores a 0", func(t *testing.T) {
		_, err := promedNotas([]float64{-20.00, 0.00, -1.00})

		if err == nil {
			t.Errorf("Las notas deben ser mayores o iguales a 0.00")
		}
	})

	t.Run("Probar error cuando se ingresa nota inferior a 0", func(t *testing.T) {
		_, err := promedNotas([]float64{-40.321})

		if err == nil {
			t.Errorf("La nota debe ser mayor o igual a 0.00")
		}
	})

}
