package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//".\veterinaria.go"
)

var _ = Describe("Listar todas las mascotas con sus datos", func() {
	var listaMasc []Mascota

	BeforeSuite(func() {
		listaMasc = append(listaMasc, Mascota{"A001", "Lassie", Animal{"perro", 33.0, 43.0}})
		listaMasc = append(listaMasc, Mascota{"B002", "Pancho", Animal{"gato", 42.0, 18.0}})
		listaMasc = append(listaMasc, Mascota{"C003", "Sebastián", Animal{"conejo", 15.0, 14.0}})
		listaMasc = append(listaMasc, Mascota{"D004", "Tony", Animal{"perro", 52.0, 13.0}})
	})

	Context("Se solicita listar registro que contiene 4 mascotas", func() {
		BeforeEach(func() {
			listaMasc
		})

		It("Se listan 4 registros de mascotas", func() {
			listaMasc.Listar(listaMasc)
			//Ω(listaMasc).Should(listaMasc)
			Expect(listaMasc).To(Equal(listaMasc))
		})
	})

})
