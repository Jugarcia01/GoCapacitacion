// Generar un aleatorio n 10 a 30
// Generar n gorutine
// Por cada gorutine, hacer un sleep de m seg, donde m es un aleatorio de 1 a 5.
// Imprimir por cada gorutina "Soy la rutina {n}"

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n, m int
	n = randInt(10, 30)
	fmt.Printf("Num de Gorutinas a lanzar: %d\n", n)
	fmt.Println("Inicio")
	for i := 1; i <= n; i++ {
		//m = randInt(1, 5)
		//time.Sleep(time.Duration(m) * time.Second)
		go func(num int) {
			m = randInt(1, 5)
			time.Sleep(time.Duration(m) * time.Second)
			fmt.Printf("sleep: %d\n", m)
			fmt.Printf("Soy la Go routine # %d\n", num)
		}(i)
	}
	fmt.Println("Fin")
	// var espera string
	// fmt.Scanln(&espera)
	time.Sleep(5 * time.Second)

}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
