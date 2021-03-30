// Autos: Marca, modelo, año, precio, *puertas, color, *transmisión (Tipo, Estado, Impuesto)
// Motos: Marca, modelo, año, *CC, precio, color (Tipo, Estado, Impuesto)

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Vehiculo struct {
	Marca, Color, Modelo, Tipo, Estado string
	Anio, Precio                       int
	Cualidad                           map[string]string
	Impuesto                           float32
}

type Carro struct {
	Vehiculo
	Puertas     int
	Transmision string
}

type Moto struct {
	Vehiculo
	CC int
}

type Accion interface {
	Encender() string
	VerInfo() string
}

// 5. Crear comportamiento Encender() para cualquier vehículo
func (v *Vehiculo) Encender() string {
	fmt.Println("Intentando encendido de:", v.Tipo, v.Marca, v.Modelo)

	switch v.Estado {
	case "Ok":
		if v.Tipo == "Moto" {
			return "La " + v.Tipo + " con capacidad de " + v.Cualidad["CC"] + " fue encendida correctamente."
		} else {
			return "El " + v.Tipo + " con transmision " + v.Cualidad["Transmision"] + " prendio exitosamente."
		}
	case "averiado":
		return "No es posible encender el vehiculo: " + v.Tipo + v.Marca + v.Modelo + " porque esta " + v.Estado
	default:
		return "Lo que se intenta encender no es un vehiculo."
	}
}

//6. Función para calcular impuesto del vehículo (46.5%) y guardarlo en el mismo objeto
func (v *Vehiculo) impuestoVehiculo() float32 {
	var tax float32
	tax = float32(v.Precio) * (46.5 / 100)
	//fmt.Printf("Impuesto del vehiculo: %0.2f", tax)
	v.Impuesto = tax
	return v.Impuesto
}

func (v *Vehiculo) estadoVehiculo() string {
	v.Estado = estadoAleatorio()
	return v.Estado
}

var err error

func main() {

	var vehiculos []Vehiculo
	vehiculos = leerArchivoVehiculos()

	//// Listando la totalidad de vehiculos
	// for _, v := range vehiculos {
	// 	fmt.Println(v)
	// }

	//// 3a. Listado de vehiculos tipo: Moto
	// var motos []Vehiculo
	// motos = listaPorTipoVehiculo(vehiculos, "Moto")
	// fmt.Println(motos)

	//// 3b. Listado de vehiculos tipo: Carro
	// var carros []Vehiculo
	// carros = listaPorTipoVehiculo(vehiculos, "Carro")
	// fmt.Println(carros)

	// Verificando impuesto de un vehiculo
	miCarro := vehiculos[randInt(0, 40)]
	fmt.Println(miCarro.Impuesto)

	fmt.Println(miCarro.Encender())

}

func leerArchivoVehiculos() []Vehiculo {
	//func leerArchivoVehiculos() {

	archivo, err := os.Open("./vehiculos.txt")

	defer func() { // Defer solo ejecuta una sentencia o una funcion. Y se auto ejecuta
		fmt.Println("El archivo fue Cerrado.")
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if err != nil {
		fmt.Println("No se pudo leer el archivo; algo salió mal...")
		panic(err)
	}
	var tipo = ""
	var vehiculos = []Vehiculo{}
	var vehic = Vehiculo{}
	Cualid := make(map[string]string)

	scanner := bufio.NewScanner(archivo)
	for i := 1; scanner.Scan(); {
		linea := scanner.Text()
		//fmt.Printf("Linea %d -> %s\n", i, linea)
		i++

		switch linea {
		case "Carros":
			tipo = "Carro"
		case "Motos":
			tipo = "Moto"
		}

		data := strings.Split(linea, ",")

		switch tipo {
		case "Carro":
			//fmt.Println("\nEn case tipo Carro")
			// fmt.Println(data)

			for ii := 0; ii < len(data); {
				if data[ii] != "Carros" {
					///fmt.Printf(data[ii] + ",")
					//Carros: Marca, modelo, año, precio, *puertas, color, *transmisión (Tipo, Estado, Impuesto)
					vehic.Marca = data[0]
					vehic.Modelo = data[1]
					vehic.Anio, err = strconv.Atoi(strings.TrimSpace(data[2]))   //vehic.Anio = data[2]
					vehic.Precio, err = strconv.Atoi(strings.TrimSpace(data[3])) //vehic.Precio = data[3]
					Cualid["Puertas"] = data[4]
					vehic.Color = data[5]
					Cualid["Transmision"] = data[6]
					vehic.Tipo = tipo
				}

				// for ind, cualidad := range Cualid {
				// 	fmt.Printf("Cualidad %s es: %s\n", ind, cualidad)
				// 	//vehic.Cualidad[ind] = cualidad
				// }

				ii++
			}

		case "Moto":
			//fmt.Println("En case tipo Moto")
			// fmt.Println(data)

			for ii := 0; ii < len(data); {
				if data[ii] != "Motos" {
					// // Motos: Marca, modelo, año, *CC, precio, color (Tipo, Estado, Impuesto)
					vehic.Marca = data[0]
					vehic.Modelo = data[1]
					vehic.Anio, _ = strconv.Atoi(strings.TrimSpace(data[2]))
					Cualid["CC"] = data[3]
					vehic.Precio, _ = strconv.Atoi(strings.TrimSpace(data[4]))
					vehic.Color = data[5]
					vehic.Tipo = tipo
				}
				ii++
			}

		}
		vehic.Cualidad = Cualid
		//fmt.Println(vehic.Cualidad)
		vehic.impuestoVehiculo()
		vehic.estadoVehiculo()

		vehiculos = append(vehiculos, vehic)
	}

	return vehiculos
}

func listaPorTipoVehiculo(vehiculos []Vehiculo, tipoVehiculo string) []Vehiculo {
	var vehics = []Vehiculo{}

	switch tipoVehiculo {
	case "Carro":
		for _, v := range vehiculos {
			if v.Tipo == tipoVehiculo {
				//fmt.Println(v)
				vehics = append(vehics, v)
			}
		}
	case "Moto":
		for _, v := range vehiculos {
			if v.Tipo == tipoVehiculo {
				//fmt.Println(v)
				vehics = append(vehics, v)
			}
		}
	case "Todos":
		for _, v := range vehiculos {
			//fmt.Println(v)
			vehics = append(vehics, v)
		}
	default:
		fmt.Println("No existe ese tipo de vehiculo:", tipoVehiculo)
	}

	return vehics

}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func estadoAleatorio() string {
	if randInt(0, 10) < 5 {
		return "averiado"
	} else {
		return "Ok"
	}
}
