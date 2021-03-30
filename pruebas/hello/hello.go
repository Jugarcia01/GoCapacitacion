package main

import "fmt"

func main() {
	fmt.Println(saludar("Julian"))
}

func saludar(t string) string {
	var s string
	if t != "" {
		s = "Hola " + t
	} else {
		s = "Hola mundo!"
	}
	return s
}

// func saludarPersona(t string) string {
// 	s := "Hola " + t
// 	return s
// }
