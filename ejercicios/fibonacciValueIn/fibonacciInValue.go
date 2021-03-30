package main

import "fmt"

// fibonacci es una función que devuelve
// una función que devuelve un int.

var a []int

func fibonacci() func(int) int {

	return func(x int) int {
		if x >= 2 {
			xf := a[x-1] + a[x-2]
			a = append(a, xf)
			return xf
		}

		a = append(a, x)
		return (a[x])
	}
}

func main() {

	var valSerie = 20

	f := fibonacci()
	for i := 0; i < valSerie; i++ {
		fmt.Println(f(i))
	}
}
