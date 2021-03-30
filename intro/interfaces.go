package main

import "fmt"

type Animal interface {
	Mover() string
	Comer(comida string) string
}

type Perro struct {
	nombre string
	patas  string
}

type Serpiente struct {
	nombre string
}

func (p Perro) Mover() string {
	return "Caminando... moviendo " + p.patas + " patas\n"
}

func (p Perro) Comer(c string) string {
	return "Esta comiendo " + c
}

func (s Serpiente) Mover() string { // Método publico porque inicia con Mayuscula
	return "Reptando...\n"
}

func (s Serpiente) Comer(c string) string {
	return "Acaba de devorar " + c
}

func iniciar(animal Animal) { // Nombre de funcion Privado inicia en minuscula
	fmt.Printf(animal.Mover())
}

func main() {
	perro := Perro{"Danger", "4"}
	anaconda := Serpiente{"Conda"}

	iniciar(perro)
	iniciar(anaconda)

	fmt.Println(perro.Comer("huesos"))
	fmt.Println(anaconda.Comer("una liebre"))

}
