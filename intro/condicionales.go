package main

import (
	"fmt"
)

func main() {

	x := 100
	y := 100
	if x > y {
		fmt.Println("x es mayor")
	} else if y > x {
		fmt.Println("y es mayor")
	} else {
		fmt.Println("y es igual a x")
	}

}
