package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	archivo, err := ioutil.ReadFile("./vehiculos.txt")

	if err != nil {
		fmt.Println("Algo salió mal")
		fmt.Println(err)
	} else {
		fmt.Printf(string(archivo))
	}
}
