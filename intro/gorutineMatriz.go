/*
	[] [] [] []
	[] [] [] []
	[] [] [] []
	3x4
Generar un n aleatorio (filas)
Generar n gorutines
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var filas = randInt(1, 5)
	const f int = 3
	const c int = 4
	mSlice := [f][c]int{}

	fmt.Println("Inicio")
	val := 0
	for fi := 0; fi < f; fi++ {
		for co := 0; co < c; co++ {
			go func(fi, co, val int) {
				val = randInt(1, 100)
				//val++
				mSlice[fi][co] = val
				fmt.Printf("[%d][%d]=%d \t", fi, co, val)
			}(fi, co, val)
		}
	}

	// Matriz secuencial
	// go func(val int) {
	// 	for fi := 0; fi < f; fi++ {
	// 		for co := 0; co < c; co++ {
	// 			//val = randInt(1, 100)
	// 			val++
	// 			mSlice[fi][co] = val
	// 			//m := randInt(1, 500)
	// 			//time.Sleep(time.Duration(m) * time.Millisecond)
	// 			fmt.Printf("[%d][%d]=%d \t", fi, co, val)
	// 		}
	// 		println()
	// 	}
	// }(val)

	fmt.Println("Fin")
	var espera string
	fmt.Scanln(&espera)
	//time.Sleep(5 * time.Second)

}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
