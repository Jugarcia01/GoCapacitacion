package main

import (
	"fmt"
)

type Hotel struct {
	// Hotel: Nombre, habitaciones, ingresos.
	//Habitacion map[string]string
	//Espacios
	Nombre       string
	Pisos        int
	Habitaciones int
	Ingresos     int
}

type Animal struct {
	tipo string
	edad float32
	peso float32
}

type Mascota struct {
	idMasc string
	nombre string
	Animal
}

type Vacuna struct {
	idVac string
	cant  float32
	Mascota
}

func main() {
	var opcion = "inicio"
	var opc = 0

	for opcion != "salir" {
		opcion = menuVeterinaria(opc)
	}

}

func menuVeterinaria(opc int) string {
	//var opc = 0
	//var habs []Habitacion
	//var verHoteles = ""
	//var opcion = 0

	fmt.Println("\nSISTEMA DE GESTIÓN DE VETERINARIA")
	fmt.Println(`Digite una opción y oprima la tecla Enter:
	1. Ver todas las mascotas
	2. Buscar
	3. Registrar mascota
	4. Dar salida
	5. Ver Vacunas
	6. Aplicar Vacuna / Jarabe
	7. Salir`)

	fmt.Scanf("%d\n", &opc)
	fmt.Printf("Opción ingresada: ")

	switch opc {
	case 1:
		fmt.Println("Ver todas las mascotas")
		// for verHoteles != "salir" {
		// 	verHoteles = listarHoteles(hoteles, habs)
		// }
		return "continuar"
	case 2:
		fmt.Println("Buscar")
		// var newHotel Hotel
		// newHotel.registrarHotel()
		// hoteles = append(hoteles, newHotel)
		return "continuar"
	case 3:
		fmt.Println("Registrar mascota")
		// calcularIngresos()
		return "continuar"
	case 4:
		fmt.Println("Dar salida")
		return "continuar"
	case 5:
		fmt.Println("Ver vacunas")
		return "continuar"
	case 6:
		fmt.Println("Aplicar Vacuna / Jarabe")
		return "continuar"
	case 7:
		fmt.Println("SALIR.\nLe deseamos un excelente día, hasta pronto!!!.")
		fmt.Println()
		return "salir"
	default:
		fmt.Println("NO EXISTE ESA OPCIÓN!!!")
	}
	return "continuar"

}
