package main

import (
	"log"	
	"net/http"
	"html/template"
)


const PUERTO = ":5000"

func main() {
	Se genera el servidor
	log.Println("Corriendo en el puerto", PUERTO)
	
	http.HandelFunc("/archivo", responderHTML)	

	log.Fatal(http.ListenAndServe(PUERTO, nil))
}

func responderHTML(w http.ResponseWriter, r *http.Request){
	plantilla, err := template.ParseFiles("./public/archivo1.html")
	if err != nil {
		log.Println("Error generando el template")
		panic(err)
	}
	plantilla.Execute(w, nil)

}