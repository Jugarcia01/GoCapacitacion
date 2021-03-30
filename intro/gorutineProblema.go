package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Estructura struct {
	mapa  map[string]int
	mutex sync.Mutex // Permite controlar el acceso/modif. de info de forma simultanea
}

var est = Estructura{mapa: make(map[string]int)}

func main() {
	wg := &sync.WaitGroup{} // Hasta que todos los procesos reporten info, espera
	wg.Add(100)
	for i := 0; i < 100; i++ {
		// guardar(fmt.Sprintf("obj_%d", rand.Intn(100)), rand.Intn(100), wg)
		go guardar(fmt.Sprintf("obj_%d", rand.Intn(100)), rand.Intn(100), wg)
	}
	go leerDato(fmt.Sprintf("obj_%d", randInt(0, 100)))
	wg.Wait() // Espera a que todas las goroutines llamen a wg.Done()
	fmt.Println(est.mapa)

}

func guardar(clave string, valor int, wg *sync.WaitGroup) {
	est.mutex.Lock()
	est.mapa[clave] = valor
	est.mutex.Unlock()
	wg.Done()
}

func leerDato(clave string) {
	est.mutex.Lock()
	fmt.Printf("Leyendo mapa en [%s] -> %d\n", clave, est.mapa[clave])
	est.mutex.Unlock()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
