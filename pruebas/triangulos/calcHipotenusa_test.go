package main

import (
	"math"
	"testing"
)

func TestCalcHipotenusa(t *testing.T) {

	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := calcHipotenusa(10.00, 20.00)
		esperado := math.Sqrt((10.00 * 10.00) + (20.00 * 20.00))

		if resultado != esperado {
			t.Errorf("Resultado : %f, se esperaba: %f", resultado, esperado)
		}
	})

	t.Run("Probar que existe error cuando la long. de catetos es negativa", func(t *testing.T) {
		_, err := calcHipotenusa(-20.00, -1.00)

		if err == nil {
			t.Errorf("Los valores de los catetos deben ser valores positivos mayores a 0")
		}
	})

	t.Run("Probar que Existe error cuando se ingresa long 0 a alguno de catetos", func(t *testing.T) {
		_, err := calcHipotenusa(0, 0.1)

		if err == nil {
			t.Errorf("Debe ingresar los valores de long. de los catetos")
		}
	})

}
