package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Generamos unos parametros a enviar desde el cliente
	params := make(map[string]string)
	params["idProducto"] = "100"
	params["idUsuario"] = "3"

	//Procedemos a construir la nuevaURL
	nuevaURL := generarUrl("http", "localhost:5000", "/parametros", params)

	// Devuelve un puntero a request y un error en caso de que exista
	r, err := http.NewRequest("GET", nuevaURL, nil)
	if err != nil {
		log.Println("Error en la petición...")
		panic(err)
	}

	// Se genera el Cliente
	cliente := &http.Client{}
	// Se efectua la peticion/request con el cliente creado
	response, errorResp := cliente.Do(r)
	if errorResp != nil {
		log.Println("Error en la respuesta...")
		panic(errorResp)
	}
	// Visualizamos la respuesta
	log.Println(response)
}

func generarUrl(protocolo, host, uri string, params map[string]string) string {
	nuevaURL, err := url.Parse(uri)
	if err != nil {
		log.Println("Error convirtiendo URL")
		panic(err)
	}

	nuevaURL.Host = host
	nuevaURL.Scheme = protocolo
	mapParametros := nuevaURL.Query()
	for k, v := range params {
		mapParametros.Add(k, v)
	}
	nuevaURL.RawQuery = mapParametros.Encode() //Espera un string similar a idProducto=1&idUsuario=104&... lo cual lo realizamos con el mapa desde el .Encode()
	strUrl := nuevaURL.String()
	log.Println("URL", strUrl)
	return strUrl
}
