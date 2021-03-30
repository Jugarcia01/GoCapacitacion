// Fuente: https://play.golang.org/p/1rNnJl0qbqO

package main

import "fmt"

// fib retorna una funcion que a su vez retorna numeros susecivos de serie Fibonacci
// fib returns a function that returns successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Las funciones son llamadas y evaluadas de izquierda a derecha // Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f(), f(), f())
}
