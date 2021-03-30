package main

import "testing"

func TestSuma(t *testing.T) {

	// assertEqual = func(resultado, esperado int, t *testing.T) {
	// 	if resultado != esperado {
	// 		t.Errorf("Resultado: %d, se esperaba: %d", resultado, esperado)
	// 	}
	// }

	// Prueba de que no debería haber error, pero existe error
	// assertNotEqual = func(resultado, esperado int, t *testing.T) {
	// 	if err != nil {
	// 		t.Errorf("No se esperaba error")
	// 	}
	// }

	t.Run("Prueba suma correcta", func(t *testing.T) {
		resultado, _ := sumar(3, 4)
		esperado := 7

		if resultado != esperado {
			t.Errorf("Resultado: %d, se esperaba: %d", resultado, esperado)
		}
	})

	t.Run("Prueba que los resultados son diferentes", func(t *testing.T) {
		resultado, _ := sumar(3, 4)
		esperado := 8

		if resultado == esperado {
			t.Errorf("La suma está mala")
		}
	})

	t.Run("Probar que hay un error", func(t *testing.T) {
		_, err := sumar(-1, 1)

		if err == nil {
			t.Error("Se esperaba un error por números menores a 0")
		}
	})
}

func TestResta(t *testing.T) {
	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := restar(10, 5)
		esperado := 5

		if resultado != esperado {
			t.Errorf("Resultado: %d, se esperaba: %d", resultado, esperado)
		}
	})

	t.Run("Probar que los resultados no coinciden", func(t *testing.T) {
		resultado, _ := restar(8, 3)
		esperado := 4

		if resultado == esperado {
			t.Errorf("Se esperaba un resultado diferente")
		}
	})

	t.Run("Probar que Y no es mayor a X", func(t *testing.T) {
		_, err := restar(5, 7)

		if err == nil {
			t.Error("Se esperaba un error por Y mayor a X")
		}
	})
}

func TestMultip(t *testing.T) {
	t.Run("Probar resultado correcto", func(t *testing.T) {
		resultado, _ := multip(2, 2)
		esperado := 4

		if resultado != esperado {
			t.Errorf("Resultado: %d se esperaba: %d", resultado, esperado)
		}
	})

	t.Run("Probar error cuando los num son menores a 0", func(t *testing.T) {
		_, err := multip(-1, -2)

		if err == nil {
			t.Errorf("Se esperaba un error por num menores a 0")
		}
	})

}

func TestDividir(t *testing.T) {
	t.Run("La division por cero debe enviar error", func(t *testing.T) {
		_, err := dividir(10, 0)

		if err == nil {
			t.Errorf("Se esperaba error por division por 0")
		}
	})

	t.Run("Probar resultado correcto", func(t *testing.T) {
		resultado, _ := dividir(4, 2)
		esperado := 2

		if resultado != esperado {
			t.Errorf("Resultado %d, se esperaba %d", resultado, esperado)
		}
	})

	t.Run("Probar error con numeros negativos", func(t *testing.T) {
		_, err := dividir(-1, -1)

		if err == nil {
			t.Errorf("Se esperaba error por num. negativos")
		}
	})

	t.Run("Probar divisiones no exactas", func(t *testing.T) {
		resultado, _ := dividir(10, 3)
		esperado := 3
		// esperado := 4

		if resultado != esperado {
			t.Errorf("Resultado: %d, se esperaba: %d", resultado, esperado)
		}
	})
}

// func assertNil(err error, t *testing.T) {
// 	if err != nil {
// 	}
// }
