package main

import "fmt"

type Humano struct { // Se crea una estructura con dos parametros.
	nombre string
	edad   int
}

type Grupo struct {
	integrantes []Humano
}


func main() {
	var persona Humano
	persona.nombre = "Julian"
	persona.edad = 38
	fmt.Println(persona)

	persona2 := Humano{nombre: "David", edad: 18}
	fmt.Println(persona2)

	persona3 := Humano{"Anny", 3}
	fmt.Println(persona3)

	/*
	personaA := Humano{"Julian", 38}
	personaB := Humano{"Pedro", 18}
	personaC := Humano{"Andrea", 21}
	myGrupo := new(Grupo)
	myGrupo = Grupo{personaA, personaB, personaC}
	fmt.Println(myGrupo)
	/*


}
