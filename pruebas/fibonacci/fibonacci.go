/* The Fibonacci Sequence can be written as a "Rule" (see Sequences and Series).
First, the terms are numbered from 0 onwards like this:
n =		0	1	2	3	4	5	6	7	8	9	10	11	12	13	14	...
xn =	0	1	1	2	3	5	8	13	21	34	55	89	144	233	377	...
fuente: https://www.mathsisfun.com/numbers/fibonacci-sequence.html */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var lim = 1

	fmt.Print("Ingrese un numero para generar la sucesion de fibonacci: ")
	fmt.Scanf("%d", &lim)

	fibo(lim)

}

func fibo(lim int) string {
	var act, ant1, ant2 int = 0, 1, 1
	var res string

	if lim > 0 {
		for k := 0; k <= lim-1; k++ {
			ant1 = ant2
			ant2 = act
			act = ant2 + ant1
			fib := act

			fmt.Print("\t", fib)
			res = strconv.Itoa(fib)
		}

	} else {
		res = "El Numero para la serie debe ser superior a 0"
	}
	return res

}
