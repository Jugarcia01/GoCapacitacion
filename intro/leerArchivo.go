/*
IMPORTANTE: Para que funcione se debe ejecutar desde la carpeta donde se encuentra la aplicacion y el archivo a ser leido.
PS C:\GoCode\src\co.com.personalsoft\GoCapacitacion> go run .\leerArchivo2.go
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	archivo, err := ioutil.ReadFile("./hola.txt")

	if err != nil {
		fmt.Println("Algo salió mal")
		fmt.Println(err)
	} else {
		fmt.Printf(string(archivo))
	}

}
