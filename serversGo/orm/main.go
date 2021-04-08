package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // Mux permite definir que métodos HTTP se van a manejar.

	// "./database"
	"./models"
	"./respuestas"
	"github.com/jugarcia01/GoCapacitacion/serversGo/orm/database"
	// "gorm.io/gorm"
	// Documentacion de gorm: https://gorm.io/docs/
)

const PORT = ":5000"

func main() {
	database.Migrar()

	router := mux.NewRouter()
	log.Println("Server escuchando en el puerto", PORT)
	router.HandleFunc("/cursos", listarCursos).Methods("GET")
	router.HandleFunc("/curso/{curso}", verCurso).Methods("GET")
	router.HandleFunc("/curso", crearCurso).Methods("POST") // En postman aplicar metodo POST para crear registro de cursos para la BD
	router.HandleFunc("/curso/{curso}", actualizarCurso).Methods("PUT")
	router.HandleFunc("/curso/{curso}", eliminarCurso).Methods("DELETE")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))

}

func listarCursos(w http.ResponseWriter, r *http.Request) {
	// conectar()
	// migrar()
	w.Header().Add("Content-Type", "application/json")
	cursos := database.GetCursos()
	respuestas.ResponderJSON(cursos, w)
	// En browser: http://localhost:5000/cursos

}

func verCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "ID no válido", -1, w) {
		return
	}
	curso := database.VerCurso(uint(idCurso))
	respuestas.ResponderJSON(curso, w)
	// En browser: http://localhost:5000/curso/1

}

func crearCurso(w http.ResponseWriter, r *http.Request) {
	var body models.Curso
	err := json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo el body!")
	curso := database.CrearCurso(body.Nombre, body.Descripcion)
	respuestas.ResponderJSON(curso, w)
}

func actualizarCurso(w http.ResponseWriter, r *http.Request) {
	// Obtenemos el id del curso a actualizar
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])

	if responderSiHayError(err, "ID no válido para actualizar!", -1, w) {
		return
	}
	var body models.Curso
	validarError(err, "Error convirtiendo el body!")
	curso := database.ActualizarCurso(uint(idCurso), body.Nombre, body.Descripcion)
	respuestas.ResponderJSON(curso, w)
}

func eliminarCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "ID no válido", -1, w) {
		return
	}
	database.EliminarCurso(uint(idCurso))
	respuestas.ResponderJSON(models.Respuesta{Mensaje: "Registro Eliminado Correctamente!", Codigo: 1}, w)
}

func validarError(err error, mensaje string) {
	if err != nil {
		log.Println(mensaje)
		panic(err)
	}
}

func responderSiHayError(err error, mensaje string, codigo int, w http.ResponseWriter) bool {
	if err != nil {
		respuestas.ResponderJSON(
			models.Respuesta{Mensaje: mensaje, Codigo: codigo}, w)
		return true
	}
	return false
}
