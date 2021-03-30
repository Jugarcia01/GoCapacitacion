// Incluyendo paquetes externos
// go get
// Referencia: https://golang.org/pkg/cmd/go/internal/get/

//Ejecutar en la terminal go env -w GO111MODULE=auto // go env -w GO111MODULE=on // go env -w GO111MODULE=off
// go mod init {nombre_paquete}
// go mod init github.com/jugarcia01/goCode/GoCapacitacion/fizzbuzz
// Permite cargar modulos a github
// Referencia: https://blog.golang.org/using-go-modules

// ejem: go get github.com/stretchr/testify
// ejem: go mod init fizzbuzz

package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	const MSG_NO_VALIDO = "No es un núm válido"
	const FIZZ = "fizz"
	const BUZZ = "buzz"
	const FIZZBUZZ = "fizzbuzz"

	t.Run("Probar cuando enviamos 0", func(t *testing.T) {
		salida := fizzbuzz(0)

		assert.Equal(t, MSG_NO_VALIDO, salida)
		// if salida != MSG_NO_VALIDO {
		// 	t.Errorf("Resultado: `%s ; se esperaba `%s", salida, MSG_NO_VALIDO)
		// }
	})

	t.Run("Probar cuando enviamos un num negativo", func(t *testing.T) {
		salida := fizzbuzz(-2)

		assert.Equal(t, MSG_NO_VALIDO, salida)
	})

	t.Run("Probar cuando enviamos un multiplo de 3", func(t *testing.T) {
		salida := fizzbuzz(3)

		assert.Equal(t, FIZZ, salida)
	})

	t.Run("Probar cuando enviamos un multiplo de 5", func(t *testing.T) {
		salida := fizzbuzz(5)

		assert.Equal(t, BUZZ, salida)
	})

	t.Run("Probar cuando enviamos un multiplo de 3 y 5", func(t *testing.T) {
		salida := fizzbuzz(15)

		assert.Equal(t, FIZZBUZZ, salida)
	})

	t.Run("Probar cuando num es positivo y No es un multiplo de 3 ni de 5", func(t *testing.T) {
		salida := fizzbuzz(22)
		esperado := strconv.Itoa(22)

		assert.Equal(t, esperado, salida)
	})

}
