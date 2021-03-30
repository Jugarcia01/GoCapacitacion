package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Inicio del main")
	ejecutarFuncion()
	fmt.Println("Fin del main")

}

func verDefer() {
	fmt.Println("Iniciando método verDefer")
	defer fmt.Println("Esto es ejecutado con Defer")
	fmt.Println("Este es el fin del metodo verDefer.")
}

func leerArchivo() {
	archivo, err := os.Open("./holaAAA.txt")

	defer func() { // Defer solo ejecuta una sentencia o una funcion. Y se auto ejecuta
		fmt.Println("Cerrando archivo...")
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if err != nil {
		fmt.Println("Algo salió mal...")
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)
	for i := 1; scanner.Scan(); {
		linea := scanner.Text()
		fmt.Printf("Linea %d -> %s\n", i, linea)
		i++
		panic(errors.New("Simulando error"))
	}
}

func ejecutarFuncion() {
	fmt.Println("Inicio de tercera funcion")
	leerArchivo()
	fmt.Println("Fin de tercera funcion")
}
