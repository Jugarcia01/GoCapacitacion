package main

import (
	"fmt"
	"strconv"
	//"syreclabs.com/go/faker"
)

type Estudiante struct {
	Nombre    string
	Edad      int
	Cursos    []string
	Notas     []float32
	PromNotas float32
}

func (e *Estudiante) promEstudiante() (PromNotas float32) {
	var sumNotas float32
	for i := 0; i < len(e.Notas); i++ {
		sumNotas = sumNotas + e.Notas[i]
	}
	e.PromNotas = (sumNotas / float32(len(e.Notas)))
	PromNotas = e.PromNotas
	return PromNotas
}

type NewEstudiante struct {
	Nombre string
	Edad   int
}

var cantEst int = 0

func main() {
	//estud1 := new(Estudiante) // Devuelve puntero
	//estud1.Nombre = "Julian"
	//estud1.Notas = []int{1, 1}
	estud1 := Estudiante{Nombre: "Julian", Edad: 18, Cursos: []string{"Biol", "Matem"}, Notas: []float32{5.0, 4.0}}
	estud2 := Estudiante{Nombre: "Andres", Edad: 17, Cursos: []string{"Mate", "Biol"}, Notas: []float32{4.0, 2.0}}
	estud3 := Estudiante{Nombre: "Diana", Edad: 15, Cursos: []string{"Soc", "Leng"}, Notas: []float32{4.2, 4.5}}

	// Promedio de Cada Estudiante:
	fmt.Println("Promedio de Cada Estudiante:")
	fmt.Printf("\nNombre: " + estud1.Nombre + ", edad: " + strconv.Itoa(estud1.Edad))
	fmt.Printf(" - Promedio Notas: %0.2f", estud1.promEstudiante())

	fmt.Printf("\nNombre: " + estud2.Nombre + ", edad: " + strconv.Itoa(estud2.Edad))
	fmt.Printf(" - Promedio Notas: %0.2f", estud2.promEstudiante())

	fmt.Printf("\nNombre: " + estud3.Nombre + ", edad: " + strconv.Itoa(estud3.Edad))
	fmt.Printf(" - Promedio Notas: %0.2f", estud3.promEstudiante())

	fmt.Printf("\n\nPromedio de Notas del Curso:\n")
	//PromEsts := PromedioCurso([]Estudiante{*estud1, estud2})
	PromEsts := PromedioCurso([]Estudiante{estud1, estud2, estud3})
	fmt.Println(PromEsts)

	fmt.Printf("\nPromedio de Edades del Curso:\n")
	PromEdades := PromedioEdad([]Estudiante{estud1, estud2, estud3})
	fmt.Println(PromEdades)

	//crearEstudianteAleatorio()

}

/*
func crearEstudianteAleatorio() (e NewEstudiante) {
	//e []NewEstudiante := []NewEstudiante{Nombre: "", Edad: 0}
	//for {
	//	i++
	//	e[i].Nombre = faker.nombre
	//	e[i].Edad = faker.numerosAleatorios(2)
	//	if i > cant {
	//		break
	//	}
	//	fmt.Printf("\nNombre: " + e[i].Nombre + ", Edad: " + strconv.Itoa(e[i].Edad))
	//	return e
	//}

	e = NewEstudiante{}
	e.Nombre = faker.FakeName()
	e.Edad = faker.RamdomInt(10, 50)

	fmt.Printf("\nNombre: " + e.Nombre + ", Edad: " + strconv.Itoa(e.Edad))
	return e

}
*/

func PromedioCurso(es []Estudiante) (promCurso float32) {
	var sumNotas float32 = 0.0
	cantEst = len(es)

	for i := 0; i < len(es); i++ {
		sumNotas = sumNotas + es[i].PromNotas
		cantEst++
	}
	promCurso = (sumNotas / float32(len(es)))
	return promCurso
}

func PromedioEdad(es []Estudiante) (promEdad float32) {
	var sumEdad float32 = 0.0
	cantEst = len(es)

	for i := 0; i < len(es); i++ {
		sumEdad = sumEdad + float32(es[i].Edad)
		cantEst++
	}
	promEdad = (sumEdad / float32(len(es)))
	return promEdad
}
