package main

import (
	"testing"
)

func TestParImpar(t *testing.T) {

	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := parImparAcotado(10)
		esperado := true

		if resultado != esperado {
			t.Errorf("Resultado : %v, se esperaba: %v", resultado, esperado)
		}
	})

	t.Run("Probar error cuando se ingresan valores enteros negativos", func(t *testing.T) {
		_, err := parImparAcotado(-77)

		if err == nil {
			t.Errorf("Valor ingresado debe ser mayor a Cero")
		}
	})

	t.Run("Probar que solo se pueden ingresar valores hasta 100.000", func(t *testing.T) {
		_, err := parImparAcotado(222333)

		if err == nil {
			t.Errorf("El valor no puede superar a Cien Mil")
		}
	})

}
