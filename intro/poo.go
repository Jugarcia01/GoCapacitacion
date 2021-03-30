package main

import (
	"fmt"
	"strconv"
)

type Persona struct {
	nombre, apellido string
	edad             int
}

type empleado struct {
	idEmpleado string
	fecha      string
}

func (e *empleado) registrarIngreso() {
	e.fecha = "17/03/2021"
}

func (p *Persona) toString() string {
	return "\nNombre: " + p.nombre + ", Apellido: " + p.apellido + ", Edad: " + strconv.Itoa(p.edad)
}

type Usuario struct {
	//persona Persona // Se pasa estruc. de forma explicita
	Persona // Se pasa la estructura pero de forma anónima.
	empleado
	email string
	pass  string
}

func (u *Usuario) imprimirDatos() string {
	return "\nNombre: " + u.nombre + ", Apellido: " + u.apellido + ", Fecha: " + u.fecha
}

func (u *Usuario) iniciarSesion() bool {
	return true
}

func main() {
	usuario := Usuario{Persona{nombre: "Julian"}, empleado{"JUGARCIA", "10/10/2020"}, "jugarcia01@gmail.com", ""}
	fmt.Println(usuario)
	fmt.Println(usuario.nombre)
	fmt.Println(usuario.idEmpleado)
	usuario.registrarIngreso()
	fmt.Println(usuario.iniciarSesion())
	fmt.Println(usuario.toString())
	fmt.Println(usuario.imprimirDatos())

}
