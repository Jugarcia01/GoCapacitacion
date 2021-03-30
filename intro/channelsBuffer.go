package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	canal := make(chan int, 1) // Los channels permiten la comunicac. entre goroutines, se crean con make y se les asigna un tipo de valor
	// canal := make(chan int, 9) // CUIDADO!!!: Un tamaño de buffer alto "9" en el canal hace que se presente comportamientos indeseados.
	salida := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2) // Este valor es importante
	// Funcion para colocar info en el canal
	go publicar(canal, salida, wg)
	// Ahora se establece mecanismo de escucha
	go escuchar(canal, salida, wg)
	wg.Wait()
}

func publicar(canal chan<- int, salida chan struct{}, wg *sync.WaitGroup) { //Con los caract de <- se restringe que el canal recibe info para publicar / Escritura en el canal
	time.Sleep(200 * time.Millisecond) // Se bloquea por 5seg
	fmt.Println("Publicando...")
	canal <- 10 // Se le publica un valor entero al canal. Canal se coloca al lado izq. para recibir el valor
	canal <- 20
	canal <- 30
	canal <- 40
	canal <- 50
	canal <- 60
	canal <- 70
	canal <- 80
	canal <- 90
	canal <- 100
	fmt.Println("Saliendo...")
	salida <- struct{}{}
	wg.Done()
}

func escuchar(canal <-chan int, salida chan struct{}, wg *sync.WaitGroup) { //Con los caract de <- se restringe que el canal reciba info para escuchar / Lectura, se va a leer los datos que hay en el canal
	fmt.Println("Escuchando...")

	ciclo := true
	for ciclo {
		select {
		case dato := <-canal:
			fmt.Println("Recibido: ", dato)
		case <-salida:
			fmt.Println("Salida.")
			ciclo = false
			//wg.Done()
			//break
		}
	}
	wg.Done()
}
