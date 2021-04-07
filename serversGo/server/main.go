package main

import (
	//Paq necesario para iniciar el servidor web
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Curso struct {
	// Se envian las propiedades publicas (Iniciales en Mayusculas) pq sino no se
	// visualizan en el browser.
	Id          int
	Nombre      string
	Descripcion string
}

const PUERTO = ":5000"

func main() {

	log.Println("Corriendo en el puerto", PUERTO)

	// Manejo de rutas, con HandleFunc con una func. que necesariamente tenga un ResponseWriter y un Request
	// en el handle se pasa la raiz, y una func. Se debe realizar para cada ruta que uno tenga.
	http.HandleFunc("/hola", hola)
	http.HandleFunc("/redirect", redireccion)
	http.HandleFunc("/not-found", noEncontrado)
	http.HandleFunc("/error", errorPeticion)
	http.HandleFunc("/parametros", parametros)
	http.HandleFunc("/enviar-json", enviarJson)

	// Para levantar un server se requiere http.ListenAndServe
	http.ListenAndServe(PUERTO, nil)

}

func hola(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Web!") // En w que es un flujo de string se escribe la cadena
	fmt.Fprint(w, "\nHola con fmt.Fprint")
}

func redireccion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirigiendo a hola")
	http.Redirect(w, r, "/hola", 301)
}

func noEncontrado(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func errorPeticion(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error en el servidor", 502)
}

func parametros(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println(r.URL.RawQuery)

	// Obteniendo los parametros en un mapa
	mapParametros := r.URL.Query()
	log.Println(mapParametros)
	log.Println(mapParametros["idProducto"][0])
	fmt.Fprintln(w, "Recibiendo parámetros")
}

func enviarJson(w http.ResponseWriter, r *http.Request) {
	// Pasando datos en formato Json
	curso1 := Curso{1, "Curso de Go", "Curso de introducción a Go/Golang"}
	curso2 := Curso{2, "Curso de Flutter", "Curso de Flutter"}
	curso3 := Curso{3, "Curso de Java", "Curso avanzado de Java"}
	cursos := []Curso{curso1, curso2, curso3}
	json.NewEncoder(w).Encode(cursos)
}
