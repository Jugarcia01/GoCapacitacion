package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

	var (
		granLibro  Books
		cortoLibro Books
	)

	BeforeEach(func() {
		granLibro = Books{
			Titulo:  "El inversor inteligente",
			Autor:   "Benjamin Graham",
			Paginas: 624,
		}

		cortoLibro = Books{
			Titulo:  "Lo que queda del dia",
			Autor:   "Kazuo Ishiguro",
			Paginas: 245,
		}
	})

	Describe("Categorizando libro por su longitud", func() {

		Context("Con mas de 300 paginas", func() {
			It("Debería ser una Novela", func() {
				Expect(granLibro.CategoryByLength()).To(Equal("NOVELA"))
			})
		})

		Context("Con menos de 300 paginas", func() {
			It("Debería ser una sintesis", func() {
				Expect(cortoLibro.CategoryByLength()).To(Equal("SINTESIS"))
			})
		})
	})

})
