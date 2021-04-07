package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "postgres"
	dbname   = "goDB"
	host     = "localhost"
	port     = "5432"
)

type Categoria struct {
	Id          int
	Nombre      string
	Descripcion string
}

var db *sql.DB
var err error

func main() {
	insertarCategoria("Cat 101", "Descripción categoria 101")
	listarCategorias()

}

func conectar() {
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//db, err := sql.Open("sqlite3", "catgorias.db")
	db, err = sql.Open("postgres", conexion)
	if err != nil {
		log.Println("Error en la conexión")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error en ping a la BD")
		panic(err)
	}

	log.Println("Conexión Abierta!")

}

func listarCategorias() {
	conectar()
	defer func() {
		db.Close()
		log.Println("Conexion Cerrada.")
	}()

	query := "SELECT id, nombre, descripcion FROM categorias"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error en la consulta")
		panic(err)
	}

	listaCat := []Categoria{}

	for rows.Next() {
		cat := Categoria{}
		// Requiere tres punteros
		rows.Scan(
			// OJO: Deben estar los campos en el mismo orden que estan en la Tabla de la BD
			&cat.Id,
			&cat.Nombre,
			&cat.Descripcion,
		)
		listaCat = append(listaCat, cat)
	}
	log.Println(listaCat)

}

func insertarCategoria(nombre, description) {
	conectar()
	defer func() {
		//
	}()

	query := fmt.Sprintf("INSERT INTO categorias(nombre, descripcion) VALUES ('%s', '%s')", nombre, descripcion)
	log.Println(query)
	_, err := db.Query(query)
	if err != nil {
		log.Println("Error insetando categoria")
		panic(err)
	}
}
