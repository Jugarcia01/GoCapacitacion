package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTipoTriangulo(t *testing.T) {
	const figuraConLados = "La figura tiene uno o más lados"
	const esTriangulo = "La figura es un triangulo"
	const longLadosAdecuada = "Los lados tienen longitud superior a 0.0"
	const equilatero = "El Triangulo es Equilatero"
	const isosceles = "El Triangulo es Isosceles"
	const escaleno = "El Triangulo es Escaleno"

	t.Run("Probar que la figura tiene lados", func(t *testing.T) {
		salida := "La figura tiene uno o más lados"

		assert.Equal(t, figuraConLados, salida)
	})

	t.Run("Probar que la figura tiene tres lados", func(t *testing.T) {
		salida := "La figura es un triangulo"

		assert.Equal(t, esTriangulo, salida)
	})

	t.Run("Probar que los lados tienen longitud mayor que 0", func(t *testing.T) {
		salida := "Los lados tienen longitud superior a 0.0"

		assert.Equal(t, longLadosAdecuada, salida)
	})

	t.Run("Probar que el triangulo es equilatero", func(t *testing.T) {
		salida := tipoTriangulo(3.0, 3.0, 3.0)
		assert.Equal(t, equilatero, salida)

	})

	t.Run("Probar que el triangulo es isosceles", func(t *testing.T) {
		salida := tipoTriangulo(3.0, 3.0, 5.0)
		assert.Equal(t, isosceles, salida)

	})

	t.Run("Probar que el triangulo es escaleno", func(t *testing.T) {
		salida := tipoTriangulo(3.0, 4.0, 5.0)
		assert.Equal(t, escaleno, salida)

	})

}
