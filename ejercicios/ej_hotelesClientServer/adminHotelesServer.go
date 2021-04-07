/*
Simular administración de una línea de hoteles por consola.
Se podrá crear:
listar hoteles (además de ver otras opciones por hotel),
crear hoteles,
calcular el total de ingresos.

Al listar hoteles, se podrá seleccionar un hotel y se visualizará estas opciones:
Ver habitaciones,
crear habitación,
reservar / liberar habitación.

Las propiedades que deben tener hoteles y habitaciones son las siguientes:
Hotel: Nombre, habitaciones, ingresos.
Habitación: Número o ID, piso, estado, precio.

INTERFAZ DE USUARIO

Linea de Hoteles
1. Ver hoteles
2. Registrar hotel
3. Calcular total ingresos
4. Salir
OPCIÓN: 1

1.Hotel Stars   2.Hotel Resort  3.Hotel Marriot   4.Regresar
OPCIÓN: 3

Hotel Marriot
1. Ver habitaciones
2. Crear habitación
3. Reservar/Liberar habitación
4. Regresar
OPCIÓN: 3

Piso1: [ Habitación 101		Habitación 102*		Habitación 103]
Piso2: [ Habitación 201		Habitación 202		Habitación 203]
Piso3: [ Habitación 301		Habitación 302		Habitación 303*]

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Hotel struct {
	// Hotel: Nombre, habitaciones, ingresos.
	//Habitacion map[string]string
	//Espacios
	Nombre       string
	Pisos        int
	Habitaciones map[int]Habitacion
	Ingresos     int
}

type Piso struct {
	// Piso: Número o ID
	//Hotel
	Id int
}

type Habitacion struct {
	// Habitación: Número o ID, piso, estado, precio.
	Id int
	//Piso
	Estado string
	Precio int
}

type accionHotel interface {
	verHotel()
	registrarHotel()
	calcularIngresos()
	salir()
}

// func (h *Hotel) verHotel() {
// 	if h.Nombre == "Nombre" {
// 		fmt.Printf("Observando informacion del hotel: %s", h.Nombre)
// 	} else {
// 		fmt.Println("El hotel digitado No existe")
// 	}
// }

func (h *Hotel) registrarHotel() *Hotel {
	fmt.Println("REGISTRA EL HOTEL")
	var newHotel Hotel
	fmt.Printf("Nombre del Hotel: ")
	fmt.Scanf("%s\n", &newHotel.Nombre)
	fmt.Printf("Ingrese Total de Pisos del Hotel: ")
	fmt.Scanf("%d\n", &newHotel.Pisos)
	fmt.Printf("Ingrese Cantidad de Habitaciones por Piso del Hotel: ")
	fmt.Scanf("%d\n", &newHotel.Habitaciones)

	newHotel.Ingresos = 0
	*h = newHotel
	return &newHotel
}

func (mapHab *Hotel) calcularIngresosHoteles() int {
	ingresos := 0
	// for num, nameH := range hoteles {
	// 	fmt.Printf("%d. %s\n", num, nameH.Nombre)
	// }
	ingresos = 120000
	return ingresos

}

type accionHabitacion interface {
	verHabitacion()
	crearHabitación()
	estadoHabitación() // Ocupar, Reservar, Liberar
	regresar()
}

var err error

// func (hab *Habitacion) registrarHabitacion() *Habitacion {
// 	fmt.Println("REGISTRO HABITACIÓN EN HOTEL: ", hab.Hotel.Nombre)
// 	var newHab Habitacion
// 	newHab.Hotel = *&hab.Hotel

// 	fmt.Println("Indique el num del piso donde desea registrar la habitación:")
// 	fmt.Scanf("%d\n", &newHab.Piso.Id)
// 	fmt.Println("Piso registrado:", newHab.Piso.Id)
// 	fmt.Println("Total Pisos del Hotel:", *&hab.Hotel.Pisos)

// 	if newHab.Piso.Id <= *&hab.Hotel.Pisos {
// 		//newHab.Piso.Hotel = *&hab.Hotel

// 		fmt.Println("Numero de la habitación que desea registrar:")
// 		fmt.Println("Digite un valor de 1 al ", *&hab.Hotel.Habitaciones, "\n")

// 		fmt.Scanf("%d\n", &newHab.Id)
// 		if newHab.Id <= *&hab.Hotel.Habitaciones {
// 			fmt.Println("Num. Habitacion: %d\n", &newHab.Id)

// 			fmt.Println("Indique el precio de la habitación:")
// 			fmt.Scanf("%d\n", &newHab.Precio)
// 			fmt.Println("Precio registrado: %d\n", &newHab.Precio)

// 			fmt.Println("Indique el estado de la habitación: (1.Libre, 2.Ocupado, 3.Reservada)")
// 			fmt.Scanf("%s\n", &newHab.Estado)
// 			fmt.Println("Precio registrado: %s\n", &newHab.Estado)

// 			hab = &newHab
// 			return &newHab
// 		} else {
// 			fmt.Println("ERROR: Numero de habitación No válido")
// 			return &newHab
// 		}
// 	}
// 	return &newHab
// }

var Hoteles = make([]Hotel, 0)

const PUERTO = ":5000"

func main() {
	log.Println("Corriendo en el puerto", PUERTO)

	// Manejo de rutas, con HandleFunc con una func. que necesariamente tenga un ResponseWriter y un Request
	// en el handle se pasa la raiz, y una func. Se debe realizar para cada ruta que uno tenga.
	loadHoteles()

	log.Println(Hoteles)

	http.HandleFunc("/listHoteles", listHoteles)
	http.HandleFunc("/registraHotel", regHotel)
	// http.HandleFunc("/ingresosHotel", ingHotel)
	// http.HandleFunc("/verHabitaciones", verHabs)
	// http.HandleFunc("/crearHabitaciones", crearHabs)
	// http.HandleFunc("/reservaHabitaciones", reservHabs)

	http.HandleFunc("/redirect", redireccion)
	http.HandleFunc("/not-found", noEncontrado)
	http.HandleFunc("/error", errorPeticion)
	http.HandleFunc("/parametros", parametros)
	http.HandleFunc("/enviar-json", enviarJson)

	// Para levantar un server se requiere http.ListenAndServe
	http.ListenAndServe(PUERTO, nil)

	// var opc = "inicio"
	// for opc != "salir" {
	// 	opc = menuPpal()
	// }
}

func listHoteles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listado de Hoteles Registrados\n")
	json.NewEncoder(w).Encode(Hoteles)
}

func regHotel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registrar Hotel\n")
	json.NewEncoder(w).Encode(Hoteles)

	// Hotel:{
	// 	Nombre       string
	// 	Pisos        int
	// 	Habitaciones map[int]Habitacion
	// 	Ingresos     int
	// }

	// Obteniendo los parametros en un mapa
	mapParam := r.URL.Query()
	log.Println(mapParam)
	hNombre := mapParam["Nombre"][0]
	hPisos, _ := strconv.Atoi(mapParam["Pisos"][1])
	hIng, _ := strconv.Atoi(mapParam["Ingresos"][2])
	hHabs := map[int]Habitacion{1: {Id: 1, Estado: "Libre", Precio: 100}}
	fmt.Fprintln(w, "Recibiendo parámetros")
	newH := Hotel{Nombre: hNombre, Pisos: hPisos, Ingresos: hIng, Habitaciones: hHabs}

	log.Println(mapParam["Nombre"][0])
	log.Println(mapParam["Pisos"][1])
	log.Println(mapParam["Ingresos"][2])

	Hoteles = append(Hoteles, newH)
	log.Println(Hoteles)

}

func redireccion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirigiendo a hola")
	http.Redirect(w, r, "/hola", 301)
}

func noEncontrado(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func errorPeticion(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error en el servidor", 502)
}

func parametros(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println(r.URL.RawQuery)

	// Obteniendo los parametros en un mapa
	mapParametros := r.URL.Query()
	log.Println(mapParametros)
	log.Println(mapParametros["idProducto"][0])
	fmt.Fprintln(w, "Recibiendo parámetros")
}

func loadHoteles() {
	hotel1 := Hotel{"Hotel Stars", 3, map[int]Habitacion{1: {Id: 02, Estado: "Libre", Precio: 200}}, 0}
	hotel2 := Hotel{"Hotel Resort", 5, map[int]Habitacion{1: {Id: 02, Estado: "Libre", Precio: 200}}, 0}
	hotel3 := Hotel{"Hotel Marriot", 10, map[int]Habitacion{2: {Id: 03, Estado: "Ocupada", Precio: 300}}, 0}
	Hoteles = append(Hoteles, hotel1)
	Hoteles = append(Hoteles, hotel2)
	Hoteles = append(Hoteles, hotel3)

}

func enviarJson(w http.ResponseWriter, r *http.Request) {
	// Pasando datos en formato Json

	//Hotel {Nombre/string, Pisos/int, Habitaciones/[int]habitaciones, Ingresos/int }
	hotel1 := Hotel{"Hotel Stars", 3, map[int]Habitacion{1: {Id: 02, Estado: "Libre", Precio: 200}}, 0}
	hotel2 := Hotel{"Hotel Resort", 5, map[int]Habitacion{1: {Id: 02, Estado: "Libre", Precio: 200}}, 0}
	hotel3 := Hotel{"Hotel Marriot", 10, map[int]Habitacion{2: {Id: 03, Estado: "Ocupada", Precio: 300}}, 0}
	hoteles := []Hotel{hotel1, hotel2, hotel3}
	json.NewEncoder(w).Encode(hoteles)
}

func menuPpal() string {
	var opc = 0
	var habs []Habitacion

	fmt.Println("\nSISTEMA DE ADMINISTRACIÓN DE HOTELES")
	fmt.Println(`Digite una opción y oprima la tecla Enter:
	1. Ver hoteles
	2. Registrar hotel
	3. Calcular total ingresos
	4. Salir`)

	fmt.Scanf("%d\n", &opc)
	fmt.Printf("Opción ingresada: ")

	var verHoteles = ""
	switch opc {
	case 1:
		fmt.Println("Ver hoteles")
		for verHoteles != "salir" {
			verHoteles = listarHoteles(Hoteles, habs)
		}
		return "continuar"
	case 2:
		fmt.Println("Registrar hotel")
		var newHotel Hotel
		newHotel.registrarHotel()
		Hoteles = append(Hoteles, newHotel)

		return "continuar"
	case 3:
		fmt.Println("Calcular total ingresos")
		// calcularIngresos()
		return "continuar"
	case 4:
		fmt.Println("SALIR.\nLe deseamos un excelente día, hasta pronto!!!.\n\n")
		return "salir"
	default:
		fmt.Println("NO EXISTE ESA OPCIÓN!!!\nDebe digitar una opción válida.")
	}
	return "continuar"
}

func listarHoteles(hoteles []Hotel, habs []Habitacion) string {
	var verHoteles = 0

	if len(hoteles) > 0 {
		fmt.Println("\nLISTADO DE HOTELES")
		for num, nameH := range hoteles {
			fmt.Printf("%d. %s\n", num, nameH.Nombre)
		}

		fmt.Printf("%d. Regresar\n", len(Hoteles))
		fmt.Println(`Digite el número del Hotel que desea visualizar y oprima la tecla Enter:`)
		fmt.Scanf("%d\n", &verHoteles)
		fmt.Println("Opción ingresada: ", verHoteles)

		switch {
		case verHoteles < len(Hoteles):
			for k, v := range Hoteles {
				if k == verHoteles {
					fmt.Println(v)
					var opcHab = "inicio"
					for opcHab != "salir" {
						opcHab = menuHabitaciones(&habs, v)
					}
				}
			}
		case verHoteles == len(Hoteles):
			fmt.Println("SALIR")
			return "salir"
		case verHoteles > len(Hoteles):
			fmt.Println("NO EXISTE ESA OPCIÓN!!!\nDebe digitar una opción válida.")
		}
		return "continuar"

	} else {
		fmt.Println("No existen hoteles registrados.")
		return "salir"
	}
}

func calculaIngresos() {
	fmt.Println("En funcion calcularTotalIngresos de Hoteles")
}

func menuHabitaciones(habs *[]Habitacion, hot Hotel) string {
	var opc = 0

	fmt.Println("\nOPCIONES PARA HABITACIONES")
	fmt.Println(`Digite una opción y oprima la tecla Enter:
	1. Ver habitaciones
	2. Crear habitación
	3. Reservar/Liberar habitación
	4. Regresar`)

	fmt.Scanf("%d\n", &opc)
	fmt.Printf("Opción ingresada: ")

	//var verHoteles = ""
	//var verHabs = ""
	switch opc {
	case 1:
		fmt.Println("Ver habitaciones")
		verHabitaciones(*habs, hot)
		//for verHabs != "salir" {
		//verHabs = verHabitaciones(habs, hot)
		//}
		return "continuar"
	case 2:
		fmt.Println("Crear habitación")
		//crearHabitacion(habs, hot)
		return "continuar"
	case 3:
		fmt.Println("Reservar/Liberar habitación")
		// calcularIngresos()
		return "continuar"
	case 4:
		fmt.Println("Regresar.\n\n")
		return "salir"
	default:
		fmt.Println("NO EXISTE ESA OPCIÓN!!!\nDebe digitar una opción válida.")
	}
	return "continuar"
}

func verHabitaciones(habs []Habitacion, hot Hotel) {
	for i := 1; i <= hot.Pisos; i++ {
		piso := strconv.Itoa(i)
		fmt.Print("Piso ", piso, ": ")

		if len(habs) != 0 {
			// for nh := 1; nh < hot.Habitaciones; {
			// 	fmt.Printf("Habitac.: ")
			// 	if i == habs[i].Piso.Id {
			// 		if nh == habs[i].Id && habs[i].Estado != "" {
			// 			fmt.Printf("%d0%d-%s \t", i, habs[i].Id, habs[i].Estado)
			// 		}
			// 	} else if habs[i].Estado == "" {
			// 		fmt.Printf("%d0#No_Activa \t")
			// 	}
			// 	nh++
			// }

			// for _, v := range habs {
			// 	if i == v.Piso.Id {
			// 		if v.Estado != "" {
			// 			fmt.Printf("%d0%d-%s \t", i, v.Id, v.Estado)
			// 		}
			// 	}
			// 	// if i == v.Piso.Id && (v.Estado != "Libre" || v.Estado != "Reserv") {
			// 	// 	fmt.Printf("%d0#-No_Activa \t", i)
			// 	// }
			// 	// fmt.Println()
			// }
		} else {
			fmt.Println("No hay habitaciones registradas.")
		}
		fmt.Println()
	}
}

// func crearHabitacion(hab *[]Habitacion, hot Hotel) {
// 	fmt.Println("REGISTRO HABITACIÓN EN HOTEL: ", hot.Nombre)
// 	var newHab Habitacion
// 	newHab.Hotel = hot

// 	fmt.Println("Indique el num del piso donde desea registrar la habitación:")
// 	fmt.Scanf("%d\n", &newHab.Piso.Id)
// 	fmt.Println("Piso registrado:", &newHab.Piso.Id)
// 	fmt.Println("Total Pisos del Hotel:", hot.Pisos)

// 	if newHab.Piso.Id <= hot.Pisos {

// 		fmt.Println("Numero de la habitación que desea registrar:")
// 		fmt.Println("Digite un valor de 1 al \n", hot.Habitaciones)

// 		// OJO Si habitacion ya existe NO la debe dejar Crear
// 		// if v.Estado == "Reserv" || v.Estado == "Libre" {
// 		// 	fmt.Println()
// 		// 	fmt.Printf("Habitación %d0%d ya existe y se encuentra en estado: %s", i, v.Id, v.Estado)

// 		// }

// 		fmt.Scanf("%d\n", &newHab.Id)
// 		if newHab.Id <= hot.Habitaciones {
// 			fmt.Println("Num. Habitacion: \n", newHab.Id)

// 			fmt.Println("Indique el precio de la habitación:")
// 			fmt.Scanf("%d\n", &newHab.Precio)
// 			fmt.Println("Precio registrado: \n", newHab.Precio)

// 			var numEst = 0
// 			var est = false
// 			for !est {
// 				fmt.Println("Indique el estado de la habitación: (1.Libre, 2.Reserv)")
// 				fmt.Scanf("%d\n", &numEst)
// 				fmt.Println("Estado: ", numEst)
// 				switch numEst {
// 				case 1:
// 					newHab.Estado = "Libre"
// 					est = true
// 				case 2:
// 					newHab.Estado = "Reserv"
// 					est = true
// 				default:
// 					fmt.Println("Estado No válido")
// 					est = false
// 				}
// 			}
// 			fmt.Println("Estado registrado: \n", newHab.Estado)
// 			*hab = append(*hab, newHab)

// 		} else {
// 			fmt.Println("ERROR: Numero de habitación No válido")
// 			err.Error()
// 		}
// 	} else {
// 		fmt.Println("ERROR: Numero de habitación No válido")
// 		err.Error()
// 	}
// }

// func estadoHabitacion() {
// }
