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
	mutex:     sync.Mutex{},
}

func main() {
	time.Sleep(1 * time.Second)
	canal := make(chan Persona, 1)
	salida := make(chan struct{})
	cantCiclos := 100
	//mapTmp := make(map[int]string)

	wg := &sync.WaitGroup{}
	wg.Add(cantCiclos)

	for i := 0; i <= cantCiclos; i++ {
		operac := intRandom(3)
		fmt.Printf("En Ciclo %d de %d - NumOperacion: %d", i, cantCiclos, operac)
		//num := intRandom(cantCiclos)
		//nameTmp := "Persona_" + strconv.Itoa(num)
		//mapTmp[i] = nameTmp

		switch operac {
		case 0:
			go imprimir(canal, salida, wg)
		case 1:
			go crear(canal, salida, wg)
		case 2:
			go editar(canal, salida, wg)
		case 3:
			go eliminar(canal, salida, wg)
		}
	}
	wg.Wait()
}

func intRandom(max int) int {
	var min int = 0
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

//func crear(canal chan<- map[int]string, salida chan struct{}, wg *sync.WaitGroup) {
func crear(canal chan<- Persona, salida chan struct{}, wg *sync.WaitGroup) {
	mapP.mutex.Lock()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Creando...")
	num := intRandom(200)
	nomb := "Persona" + strconv.Itoa(num)

	mapP.indNombre[num] = Persona.indNombre
	sliceP = append(sliceP, p)
	canal <- sliceP
	fmt.Println("Saliendo...")

	mapP.mutex.Unlock()
	salida <- struct{}{}
	wg.Done()
}

//func imprimir(canal <-chan map[int]string, salida chan struct{}, wg *sync.WaitGroup) {
func imprimir(canal <-chan Persona, salida chan struct{}, wg *sync.WaitGroup) {
	fmt.Println("Escuchando...")
	ciclo := true
	for ciclo {
		select {
		case dato := <-canal:
			fmt.Println("Recibido: ", dato)
		case <-salida:
			fmt.Println("Salida.")
			ciclo = false
		}
	}
	wg.Done()
}

//func editar(canal chan<- map[int]string, salida chan struct{}, wg *sync.WaitGroup) {
func editar(canal chan<- Persona, salida chan struct{}, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Editando...")
	num := intRandom(200)
	nomb := "Persona" + strconv.Itoa(num)
	//p := Persona{Nombre: nomb}
	sliceP[num] = Persona{Nombre: nomb}
	canal <- sliceP
	fmt.Println("Saliendo...")
	salida <- struct{}{}
	wg.Done()
}

//func eliminar(canal chan<- map[int]string, salida chan struct{}, wg *sync.WaitGroup) {
func eliminar(canal chan<- Persona, salida chan struct{}, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Eliminando...")
	num := intRandom(200)
	//nomb := "Persona" + strconv.Itoa(num)
	//p := Persona{Nombre: nomb}
	sliceP = remove(sliceP, num)
	canal <- sliceP
	fmt.Println("Saliendo...")
	salida <- struct{}{}
	wg.Done()
}

func remove(slice []Persona, s int) []Persona {
	return append(slice[:s], slice[s+1:]...)
}
