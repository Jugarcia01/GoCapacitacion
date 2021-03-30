package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Trader struct {
	id         int
	Nombre     string
	fondoValor []float32
}

type Registros struct {
	Traders map[int]Trader
	mutex   sync.Mutex
}

var mapReg = Registros{Traders: make(map[int]Trader)}

func main() {
	//----
	mapReg = iniciaTraders()
	fmt.Println(mapReg)

	canalInv := make(chan int, 1) // Los channels permiten la comunicac. entre goroutines, se crean con make y se les asigna un tipo de valor
	canalOp := make(chan float32, 1)

	// Funcion anonima que indica el numero del inversor que va a aplicar una operacion
	go func() {
		for k, v := range mapReg.Traders {
			fmt.Print("Inversor:", v.Nombre, "\t")
			canalInv <- k
		}
	}()

	// Funcion que realiza las operaciones sobre los fondos.
	go func() {
		for x := range canalInv {
			numInv := x
			for _, v := range mapReg.Traders[numInv].fondoValor {
				var multRes = intRandom(-10, 10)
				v := (float32(multRes) * 0.20 * v)
				if v >= 0 {
					canalOp <- v
				} else {
					canalOp <- 0.00
				}

			}
		}
	}()

	// Se imprime desde bucle for del main el resultado
	for _, v := range mapReg.Traders {
		for x := range canalOp {
			fmt.Printf("%0.2f \n", x)
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("%s resultado: %d", v.Nombre, v.fondoValor)
	}

}

func iniciaTraders() Registros {
	//12345, Juan,    [30000, 25000, 18000, 30000, 20000]
	mapReg.Traders[0] = Trader{id: 12345, Nombre: "Juan", fondoValor: []float32{30000.00, 25000.00, 18000.00, 30000, 20000.00}}

	//45678, Tony,    [40000, 28000, 22000, 20000, 26000]
	mapReg.Traders[1] = Trader{id: 45678, Nombre: "Tony", fondoValor: []float32{40000.00, 28000.00, 22000.00, 20000.00, 26000.00}}

	//34567, Ana,     [35000, 15000, 17000, 26000, 30000]
	mapReg.Traders[2] = Trader{id: 34567, Nombre: "Ana", fondoValor: []float32{35000.00, 15000.00, 17000.00, 26000.00, 30000.00}}

	//56789, Lina,    [21000, 35000, 20000, 18000, 28000]
	mapReg.Traders[3] = Trader{id: 56789, Nombre: "Lina", fondoValor: []float32{21000.00, 35000.00, 20000.00, 18000.00, 28000.00}}

	//23456, Ricardo, [31000, 23000, 16000, 25000, 22000]
	mapReg.Traders[4] = Trader{id: 23456, Nombre: "Ricardo", fondoValor: []float32{31000.00, 23000.00, 16000.00, 25000.00, 22000.00}}

	return mapReg
}

func intRandom(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
