package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Persona struct {
	indNombre map[int]string
	mutex     sync.Mutex
}

var mapP = Persona{
	indNombre: make(map[int]string),
}

func main() {
	//time.Sleep(1 * time.Second)
	cantCiclos := 100
	wg := &sync.WaitGroup{}
	wg.Add(cantCiclos)
	for i := 0; i <= cantCiclos; i++ {
		operac := intRandom(3)
		fmt.Printf("\nEn Ciclo %d de %d - NumOperacion: %d\t", i, cantCiclos, operac)
		clave := intRandom(cantCiclos)
		num := intRandom(cantCiclos)
		valor := "Persona_" + strconv.Itoa(num)

		switch operac {
		case 0:
			go imprimir(clave, wg)
		case 1:
			go crear(clave, valor, wg)
		case 2:
			go editar(clave, valor, wg)
		case 3:
			go eliminar(clave, wg)
		}
	}
	wg.Wait()

	mapP.mutex.Lock()
	time.Sleep(3 * time.Second)
	fmt.Println("\nTOTAL DE PERSONAS REGISTRADAS")
	fmt.Println(mapP)
	mapP.mutex.Unlock()
}

func intRandom(max int) int {
	var min int = 0
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func crear(clave int, valor string, wg *sync.WaitGroup) {
	mapP.mutex.Lock()
	fmt.Println("Creando...")
	mapP.indNombre[clave] = valor
	fmt.Printf("\nCreado registro [%d] -> %s\n", clave, mapP.indNombre[clave])
	mapP.mutex.Unlock()
	wg.Done()
}

func imprimir(clave int, wg *sync.WaitGroup) {
	mapP.mutex.Lock()
	fmt.Println("Imprimiendo...")
	_, claveExiste := mapP.indNombre[clave]
	if claveExiste {
		fmt.Printf("Imprimiendo mapa en [%d] -> %s\n", clave, mapP.indNombre[clave])
	} else {
		fmt.Println("No existe registro de persona:", clave)
	}
	mapP.mutex.Unlock()
	wg.Done()
}

func editar(clave int, valor string, wg *sync.WaitGroup) {
	mapP.mutex.Lock()
	fmt.Println("Editando...")
	_, claveExiste := mapP.indNombre[clave]
	if claveExiste {
		mapP.indNombre[clave] = valor
		fmt.Printf("Editando registro en [%d] -> %s\n", clave, mapP.indNombre[clave])
	} else {
		fmt.Println("No existe registro de persona:", clave)
	}
	mapP.mutex.Unlock()
	wg.Done()
}

func eliminar(clave int, wg *sync.WaitGroup) {
	mapP.mutex.Lock()
	fmt.Println("Eliminando...")
	_, claveExiste := mapP.indNombre[clave]
	if claveExiste {
		fmt.Printf("Eliminado registro en [%d] -> %s\n", clave, mapP.indNombre[clave])
		delete(mapP.indNombre, clave)
	} else {
		fmt.Println("No existe registro de persona:", clave)
	}
	mapP.mutex.Unlock()
	wg.Done()
}
