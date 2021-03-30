/*
Considere 5 personas que tienen algunos fondos (5) para invertir.
Se debe lanzar una go routine por cada persona
para empezar a calcular el resultado de estas inversiones.
Por cada fondo se lanza una go routine,
se obtiene el resultado de la operación y se devuelve a la rutina anterior a través de un channel
para actualizar la información de los fondos.
Para el cálculo del resultado, se generará un aleatorio entre -10 y 10
y se multiplica por el valor enviado.

ID, Nombre,   [fondo1, fondo2, fondo3, fondo4, fondo5]
12345, Juan,    [30000, 25000, 18000, 30000, 20000]
45678, Tony,    [40000, 28000, 22000, 20000, 26000]
34567, Ana,     [35000, 15000, 17000, 26000, 30000]
56789, Lina,    [21000, 35000, 20000, 18000, 28000]
23456, Ricardo, [31000, 23000, 16000, 25000, 22000]
*/

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
	mutex      sync.Mutex
}

var mapT = make(map[int]Trader)

func main() {
	cantOperaciones := 100
	canalInv := make(chan int, 1) // Los channels permiten la comunicac. entre goroutines, se crean con make y se les asigna un tipo de valor
	canalRes := make(chan float32, 1)
	salida := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(cantOperaciones)

	mapT = iniciaTraders()

	for i := 0; i <= cantOperaciones; i++ {
		go inversor(canalInv, salida, wg)
		go operacion(canalRes, salida, wg)
	}
	wg.Wait()

	mapT.mutex.Lock()
	time.Sleep(3 * time.Second)
	fmt.Println("\nTOTAL DE PERSONAS REGISTRADAS")
	fmt.Println(mapT)
	mapT.mutex.Unlock()
}

func intRandom(min int, max int) int {
	//var min int = 0
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func inversor(canalInv chan<- int, salida chan struct{}, wg *sync.WaitGroup) {
	persRandom := intRandom(0, 4)
	mapT[canalInv].mutex.Lock()
	fmt.Println("Inicia Inversor...")
	canalInv <- persRandom // Se le publica un valor entero al canal. Canal se coloca al lado izq. para recibir el valor
	fmt.Println("Sale Inversor.")
	salida <- struct{}{}

	//mapT.mutex.Unlock()
	wg.Done()
}

func operacion(canal <-chan float32, salida chan struct{}, wg *sync.WaitGroup) { //Con los caract de <- se restringe que el canal reciba info para escuchar / Lectura, se va a leer los datos que hay en el canal
	mapT.mutex.Lock()
	fmt.Println("Operando...")
	//fondoRandom := intRandom(1, 5)

	for i := 0; i < 5; i++ {
		multRes := intRandom(-10, 10)
		canal.fondoValor[i] = canal.fondoValor[i] * multRes

		select {
		case dato:
			<-canal
			fmt.Println("Recibido: ", dato)
		case <-salida:
			fmt.Println("Salida de Operacion.")
			ciclo = false
		}
	}

	ciclo := true
	for ciclo {
		select {
		case dato := <-canal:
			fmt.Println("Recibido: ", dato)
		case <-salida:
			fmt.Println("Salida de Operacion.")
			ciclo = false
			//wg.Done()
			//break
		}
	}

	mapT.mutex.Unlock()
	wg.Done()
}

// func operacWinLoss(clave int, resultado float32, wg *sync.WaitGroup) {
// 	mapP.mutex.Lock()
// 	fmt.Println("Creando...")
// 	mapP.indNombre[clave] = valor
// 	fmt.Printf("\nCreado registro [%d] -> %s\n", clave, mapP.indNombre[clave])
// 	mapP.mutex.Unlock()
// 	wg.Done()
// }

func iniciaTraders() map[int]Trader {
	mapFV := make(map[string]int)
	//mapFV = {"fondo1":40000, "fondo2":28000; "fondo3":22000; "fondo4":20000; "fondo5":26000}

	//12345, Juan,    [30000, 25000, 18000, 30000, 20000]
	mapFV["fondo1"] = 30000
	mapFV["fondo2"] = 25000
	mapFV["fondo3"] = 18000
	mapFV["fondo4"] = 30000
	mapFV["fondo5"] = 20000

	mapT[0] = Trader{
		id:         12345,
		Nombre:     "Juan",
		fondoValor: mapFV,
	}

	//45678, Tony,    [40000, 28000, 22000, 20000, 26000]
	mapFV["fondo1"] = 40000
	mapFV["fondo2"] = 28000
	mapFV["fondo3"] = 22000
	mapFV["fondo4"] = 20000
	mapFV["fondo5"] = 26000

	mapT[1] = Trader{
		id:         45678,
		Nombre:     "Tony",
		fondoValor: mapFV,
	}

	//34567, Ana,     [35000, 15000, 17000, 26000, 30000]
	mapFV["fondo1"] = 35000
	mapFV["fondo2"] = 15000
	mapFV["fondo3"] = 17000
	mapFV["fondo4"] = 26000
	mapFV["fondo5"] = 30000

	mapT[2] = Trader{
		id:         34567,
		Nombre:     "Ana",
		fondoValor: mapFV,
	}

	//56789, Lina,    [21000, 35000, 20000, 18000, 28000]
	mapFV["fondo1"] = 21000
	mapFV["fondo2"] = 35000
	mapFV["fondo3"] = 20000
	mapFV["fondo4"] = 18000
	mapFV["fondo5"] = 28000

	mapT[3] = Trader{
		id:         56789,
		Nombre:     "Lina",
		fondoValor: mapFV,
	}

	//23456, Ricardo, [31000, 23000, 16000, 25000, 22000]
	mapFV["fondo1"] = 31000
	mapFV["fondo2"] = 23000
	mapFV["fondo3"] = 16000
	mapFV["fondo4"] = 25000
	mapFV["fondo5"] = 22000

	mapT[4] = Trader{
		id:         23456,
		Nombre:     "Ricardo",
		fondoValor: mapFV,
	}

	return mapT

}
