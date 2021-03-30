package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	const MSG_NO_VALIDO = "El Numero para la serie debe ser superior a 0"

	t.Run("Probar cuando enviamos valor 0", func(t *testing.T) {
		salida := fibo(-11)
		assert.Equal(t, MSG_NO_VALIDO, salida)
	})

	t.Run("Probar cuando enviamos valor superior a 0", func(t *testing.T) {
		salida := fibo(10)
		esperado := strconv.Itoa(55)
		assert.Equal(t, esperado, salida)
	})

	t.Run("Probar que valor actual es igual a la suma de los dos anteriores", func(t *testing.T) {
		ant1 := 13
		ant2 := 21
		act := fibo(9)

		esperado := strconv.Itoa(ant1 + ant2)
		assert.Equal(t, esperado, act)
	})

}
