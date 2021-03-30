package main

import (
	"math"
	"testing"
)

func TestPerimetro(t *testing.T) {
	t.Run("Prueba con valor correcto", func(t *testing.T) {
		resultado, _ := perimetro(2.1)
		esperado := 2.00 * math.Pi * 2.10

		if math.Trunc(resultado) != math.Trunc(esperado) {
			t.Errorf("Resultado: %0.2f, se esperaba: %0.2f", resultado, esperado)
		}
	})

	t.Run("Prueba con valor de radio entero", func(t *testing.T) {
		resultado, _ := perimetro(4)
		esperado := 2.00 * math.Pi * 4

		if math.Trunc(resultado) != math.Trunc(esperado) {
			t.Errorf("Resultado: %0.2f, se esperaba: %0.2f", resultado, esperado)
		}
	})

	t.Run("Prueba error cuando valor radio negativo", func(t *testing.T) {
		_, err := perimetro(-5)

		if err == nil {
			t.Errorf("Se esperaba un error por números menores a 0")
		}
	})

	// t.Run("Prueba error cuando se ingresa texto en lugar de numero", func(t *testing.T) {
	// 	_, err := perimetro("Un Texto")

	// 	if err == nil {
	// 		t.Errorf("Se esperaba Error por ingresar texto en lugar de número.")
	// 	}

	// 	// CUIDADO con establecer mal la lógica de comparación
	// 	// if err != nil {
	// 	// 	t.Errorf("Se esperaba Error por ingresar texto en lugar de número.")
	// 	// }
	// })
}

func TestArea(t *testing.T) {
	t.Run("Prueba con valor correcto", func(t *testing.T) {
		resultado, _ := area(4.4)
		esperado := math.Pi * 4.4 * 4.4

		if math.Trunc(resultado) != math.Trunc(esperado) {
			t.Errorf("Resultado: %0.2f, se esperaba: %0.2f", resultado, esperado)
		}
	})

	t.Run("Prueba con valor de radio entero", func(t *testing.T) {
		resultado, _ := area(10)
		esperado := math.Pi * 10 * 10

		if math.Trunc(resultado) != math.Trunc(esperado) {
			t.Errorf("Resultado: %0.2f, se esperaba: %0.2f", resultado, esperado)
		}
	})

	t.Run("Prueba error cuando valor radio negativo", func(t *testing.T) {
		_, err := area(-5)

		if err == nil {
			t.Errorf("Se esperaba un error por números menores a 0")
		}
	})
}
