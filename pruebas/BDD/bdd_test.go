/*
Tener en cuenta que para que la suite de ginkgo y gomega funcionen bien, se debe
aplicar previamente los siguientes comandos:
ginkgo generate
ginkgo bootstrap

Ejem:
PS C:\GoCode\src\GoCapacitacion\pruebas\BDD> ginkgo generate
Generating ginkgo test for bdd in:
  bdd_test.go

PS C:\GoCode\src\GoCapacitacion\pruebas\BDD> ginkgo bootstrap
Generating ginkgo test suite bootstrap for bdd in:
    bdd_suite_test.go
*/

package main

import (
	. "github.com/onsi/ginkgo" // BDD
	. "github.com/onsi/gomega" // Asserts
)

var _ = Describe("Consignar un monto a una cuenta", func() {
	var iCuenta Cuenta

	BeforeSuite(func() {
		iCuenta = Cuenta{Monto: 10000.0}
	})

	Context("Si el valor a consignar es 2000.0", func() {
		BeforeEach(func() {
			iCuenta.Monto = 10000.0
		})

		It("El monto de la cuenta debe se de 12000.0", func() {
			iCuenta.Consignar(2000.0)
			Ω(iCuenta.Monto).Should(Equal(float32(12000.0)))
			Expect(iCuenta.Monto).To(Equal(float32(12000.0)))
		})
	})

	// Context("Si el valor a consignar es -2000.0", func() {
	// 	BeforeEach(func() {
	// 		iCuenta.Monto = 10000.0
	// 	})
	// 	It("El monto de la cuenta debe se de 10000.0", func() {
	// 		_, err := iCuenta.Consignar(-2000.0)
	// 		Ω(iCuenta.Monto).Should(Equal(10000.0))
	// 		Expect(err.Error()).To(Equal("El valor no debe ser negativo"))
	// 	})
	// })
})
