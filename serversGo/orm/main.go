package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Mux permite definir que métodos HTTP se van a manejar.
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// Documentacion de gorm: https://gorm.io/docs/
	//"gorm.io/driver/postgres"
)

type Curso struct {
	gorm.Model
	Nombre      string
	Descripcion string
}

type Persona struct {
	gorm.Model
	Nombre string
	Email  string
}

const PORT = ":5000"

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cursos", listarCursos).Methods("GET")
	router.HandleFunc("/curso/{curso}", verCurso).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))

}

func listarCursos(w http.ResponseWriter, r *http.Request) {
	conectar()
	migrar()
	fmt.Fprint(w, "Ruta: /cursos")
	// En browser: http://localhost:5000/cursos

}

func verCurso(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ruta: con /curso/#")
	// En browser: http://localhost:5000/curso/1

}

// Nos conectamos a la BD antes de crear tablas.
func conectar() {

	// Si la conexion a la db ya existe entonces no es necesario que se vuelva
	// a conectar. se aplica return si db es dif. de nil es decir cuando no hay conexion
	// ejecuta la conexion
	if db != nil {
		return
	}
	// Se indica a gorm que abra la conexion y donde lo va a guardar la config de la conexion
	db, err := gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})

	// Con postgres la cadena de conexion sería similar a:
	//conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//db, err = sql.Open(postgres.Open(conexion), &gorm.Config{})

	if err != nil {
		log.Println("Error en la conexion...")
		panic(err)
	}
}

// Se procede a crear la tabla mediante el programa, el ORM permite efectuar ese proceso
// Una vez se establece la conexion se pueden migrar los datos en la db.
func migrar() {
	db.AutoMigrate()
}
