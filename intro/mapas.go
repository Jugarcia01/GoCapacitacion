package main

import "fmt"

func main() {
	// Declarando e Inicializ. un mapa en la variable monedas
	monedas := make(map[string]string)

	monedas["Colombia"] = "Peso Colombiano"
	monedas["EEUU"] = "Dolar"
	monedas["España"] = "Euro"
	fmt.Println(monedas)

	//delete(monedas, "España") // Eliminando pares llave,valor en el mapa.
	//fmt.Println(monedas)

	// Recorriendo un mapa.
	for key, value := range monedas {
		fmt.Println("Clave: ", key, "   \tValor: ", value)
	}

	arr := []int{10, 20, 30, 40, 50}
	// for indice, elemento := rango arreglo{}
	for _, el := range arr {
		fmt.Println("Elemento: ", el)
	}

}
