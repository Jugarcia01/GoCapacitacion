package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	time.Sleep(4 * time.Second)
	canal := make(chan int) // Los channels permiten la comunicac. entre goroutines, se crean con make y se les asigna un tipo de valor
	wg := &sync.WaitGroup{}
	wg.Add(2) // Este valor es importante
	// Funcion para colocar info en el canal
	go publicar(canal, wg)
	// Ahora se establece mecanismo de escucha
	go escuchar(canal, wg)
	wg.Wait()
}

// func publicar(canal chan int, wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second) // Se bloquea por 5seg
// 	fmt.Println("Publicando...")
// 	canal <- 20 // Se le publica un valor entero al canal. Canal se coloca al lado izq. para recibir el valor
// 	wg.Done()
// }

// func escuchar(canal chan int, wg *sync.WaitGroup) {
// 	fmt.Println("Escuchando...")
// 	//dato := <-canal // Para obtener el valor del canal se pasa al lado derecho
// 	dato := <-canal // Se queda escuchando hasta que haya datos/algo en el canal
// 	fmt.Println("Recibido:", dato)
// 	wg.Done()
// }

func publicar(canal chan<- int, wg *sync.WaitGroup) { //Con los caract de <- se restringe que el canal recibe info para publicar / Escritura en el canal
	time.Sleep(1 * time.Second) // Se bloquea por 5seg
	fmt.Println("Publicando...")
	canal <- 20 // Se le publica un valor entero al canal. Canal se coloca al lado izq. para recibir el valor
	wg.Done()
}

func escuchar(canal <-chan int, wg *sync.WaitGroup) { //Con los caract de <- se restringe que el canal reciba info para escuchar / Lectura, se va a leer los datos que hay en el canal
	fmt.Println("Escuchando...")
	//dato := <-canal // Para obtener el valor del canal se pasa al lado derecho
	dato := <-canal // Se queda escuchando hasta que haya datos/algo en el canal
	fmt.Println("Recibido:", dato)
	wg.Done()
}
