package main

import "testing"

func TestSaludar(t *testing.T) {
	t.Run("Probar cuando el nombre esta vacio", func(t *testing.T) {
		saludo := saludar("")
		esperado := "Hola mundo!"

		if saludo != esperado {
			t.Errorf("Obtuvimos %s pero esperabamos %s", saludo, esperado)
		}
	})

	t.Run("Probar cuando se ingresa otro nombre", func(t *testing.T) {
		saludo := saludar("Julian")
		esperado := "Hola Julian"

		if saludo != esperado {
			t.Errorf("Obtuvimos %s pero esperabamos %s", saludo, esperado)
		}
	})
}

// func TestSaludarPersona(t *testing.T) {
// 	saludo := saludar("Julian")
// 	esperado := "Hola Julian"
// 	if saludo != esperado {
// 		t.Errorf("Obtuvimos %s pero esperabamos %s", saludo, esperado)
// 	}
//}
