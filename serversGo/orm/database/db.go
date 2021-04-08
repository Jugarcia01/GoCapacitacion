package database

import (
	"fmt"
	"log"

	"../models"

	"gorm.io/gorm"
	// Para obtener el driver de sqlite en Go digitar en consola: go get gorm.io/driver/sqlite
	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	// "github.com/jugarcia01/GoCapacitacion/serversGo/orm/models"
	//slog "github.com/jugarcia01/GoCapacitacion/serversGo/orm/models"
	//alog "github.com/jugarcia01/GoCapacitacion/serversGo/orm/models"
)

// Constantes de conexion a la BD
const (
	user     = "postgres"
	password = "postgres"
	dbname   = "goDB"
	host     = "localhost"
	port     = "5432"
)

var db *gorm.DB
var err error

// GetConn realiza el proceso de establecer conexion a la BD
func GetConn() *gorm.DB {
	if db != nil {
		return db
	}

	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(conexion), &gorm.Config{})
	//db, err = gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		log.Println("Error en la conexión...")
		panic(err)
	}
	Migrar()
	return db
}

// Migrar crea la tabla en la BD con sus propiedades basada en el modelo indicado
// Se procede a crear la tabla mediante el programa, el ORM permite efectuar ese proceso
// Una vez se establece la conexion se pueden migrar los datos a la db.
func Migrar() {
	//_ = GetConn()
	db.AutoMigrate(&models.Curso{})
}

// GetCursos obtiene el listado de cursos que existen en la BD
// Efectuar request GET en postman: http://localhost:5000/cursos
func GetCursos() (listarCursos []models.Curso) {
	_ = GetConn()
	db.Find(&listarCursos)
	return
}

// CrearCurso crea un registro de curso en la BD
// Efectuar el request POST en postman: http://localhost:5000/curso
func CrearCurso(nombre, descripcion string) (curso models.Curso) {
	_ = GetConn()
	curso = models.Curso{Nombre: nombre, Descripcion: descripcion}
	db.Create(&curso)
	db.Last(&curso)
	return
}

// VerCurso permite ver un registro de curso
func VerCurso(id uint) (curso models.Curso) {
	_ = GetConn()
	db.Find(&curso, id)
	return
}

// ActualizarCurso realiza la actualización de un registro de curso en la BD
func ActualizarCurso(id uint, nombre, descripcion string) (curso models.Curso) {
	_ = GetConn()
	db.Find(&curso, id)

	// //1ra. Una forma de actualizar los datos, en esta se pueden actualizar algunos datos
	// db.Model(&curso).Update("Nombre", nombre)
	// db.Model(&curso).Update("Descripcion", descripcion)

	// //2da. Para esta siguiente forma deben enviarse todos los datos.
	// db.Model(&curso).Updates(models.Curso{Nombre: nombre, Descripcion: descripcion})

	// //3ra. Permite actualizar algunas/todas las propiedades a la vez por medio de un mapa
	db.Model(&curso).Updates(map[string]interface{}{"Nombre": nombre, "Descripcion": descripcion})
	return
}

// EliminarCurso borra el registro con el id indicado en la BD
func EliminarCurso(id uint) {
	_ = GetConn()
	db.Delete(&models.Curso{}, id)
}
