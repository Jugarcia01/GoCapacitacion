package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMenuVeterinaria(t *testing.T) {
	const badOpc = "continuar"
	const exitApp = "salir"

	t.Run("Probar cuando se envia opcion de menú válida", func(t *testing.T) {
		salida := menuVeterinaria(111)

		assert.Equal(t, badOpc, salida)
	})

	t.Run("Probar cuando se desea salir de la aplicación", func(t *testing.T) {
		salida := menuVeterinaria(7)

		assert.Equal(t, exitApp, salida)
	})

	t.Run("Probar cuando se envia una opción valida del menú", func(t *testing.T) {
		salida := menuVeterinaria(3)

		assert.Equal(t, "continuar", salida)
	})

}

func TestListadoMascotas(t *testing.T) {
	const  = "continuar"
	const exitApp = "salir"

	t.Run("Probar cuando se envia opcion de menú válida", func(t *testing.T) {
		salida := menuVeterinaria(111)
		assert.Equal(t, badOpc, salida)
	})


	
}

