/*
Simular la captura de pedido de N productos por consola.
* Se mostrará el menú de los productos disponibles y la opción de confirmar pedido.
* Se solicitará que ingrese el número correspondiente al producto deseado.
* Se agrega el producto seleccionado a una lista de pedido y se regresará al menú de los productos y la opción para confirmar.
* Si selecciona la opción de confirmar y no hay productos seleccionados, mostrar el mensaje por consola y terminar el programa.

* Al confirmar el pedido se lanzarán go routines para verificar que se cuente con los elementos necesarios para preparar cada producto (NO todo el pedido).
* Si encuentra que los elementos no son suficientes para preparar algún producto, enviar por un canal la información y mostrar el mensaje por consola y terminar el programa.
* Si se tienen los elementos necesarios, lanzar go routines para "preparar" cada producto con un delay según el tiempo de cada producto. Aquí se realizará de nuevo la validación de elementos.
* Al terminar el producto, enviar por el canal la respuesta con el producto.
* Si preparando algún producto encuentra que los elementos no son suficientes, enviar la respuesta por el canal, mostrar el mensaje por consola y terminar  el programa.
* Al finalizar todos los productos, mostrar el mensaje de "Pedido terminado" y el total a pagar.

Para la comunicación entre canales, usar la estructura Respuesta, la cual tiene un entero para validar el estado, un mensaje en casos de error y un producto en caso de haber preparado alguno.

Adjunto código inicial. Si desean pueden agregar más elementos y productos
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Ingrediente struct {
	ingredientes [7]Elemento
	mutex        sync.Mutex
}

type Elemento struct {
	id       int
	nombre   string
	cantidad int
}

type Producto struct {
	id        int
	nombre    string
	precio    float32
	tiempo    int
	elementos map[int]int
}

type Respuesta struct {
	ok       int
	mensaje  string
	producto Producto
}

var carne = Elemento{1, "Carne", 5}
var pan = Elemento{2, "Pan", 10}
var salchicha = Elemento{3, "Salchicha", 3}
var tomate = Elemento{4, "Tomate", 5}
var lechuga = Elemento{5, "Lechuga", 5}
var papa = Elemento{6, "Papa", 10}
var arepa = Elemento{7, "Arepa", 3}

var hamburguesa = Producto{
	id:     1,
	nombre: "Hamburguesa",
	precio: 10000,
	tiempo: 10,
	elementos: map[int]int{
		1: 1,
		2: 2,
		4: 1,
		5: 1,
		6: 3,
	},
}

var perro = Producto{
	id:     2,
	nombre: "Perro Caliente",
	precio: 7000,
	tiempo: 6,
	elementos: map[int]int{
		3: 1,
		2: 1,
		4: 1,
		5: 1,
		6: 3,
	},
}

var arepaB = Producto{
	id:     3,
	nombre: "Arepa burger",
	precio: 8000,
	tiempo: 7,
	elementos: map[int]int{
		1: 1,
		4: 1,
		5: 1,
		6: 3,
	},
}

var elementos = [7]Elemento{carne, pan, salchicha, tomate, lechuga, papa, arepa}
var productos = [3]Producto{hamburguesa, perro, arepaB}
var ingredientes = Ingrediente{ingredientes: elementos}
var total float32

func main() {
	var pedido []Producto
	var resp []Respuesta
	ch := make(chan Respuesta, 1) //Canal para verific. existencia de Elementos

	pedido = menuProductos()
	fmt.Println(pedido)

	// Se verifica disponibilidad de ingredientes por producto en el pedido.
	for i := range pedido {
		go ProcDispByProd(pedido[i], ch)

	}

	chPrepair := make(chan Respuesta, 1) //Canal para la Preparacion de productos
	wgPrepair := &sync.WaitGroup{}
	wgPrepair.Add(1)
	// Se efectua proceso de preparación de productos.
	for i, _ := range pedido {
		go PrepairProduct(pedido[i], chPrepair, wgPrepair)
		resp = append(resp, <-chPrepair)
	}
	wgPrepair.Done()

	fmt.Println("Pedido terminado!!!")
	fmt.Printf("Total a Pagar: $%0.2f", total)

}

func menuProductos() []Producto {
	fmt.Println("Seleccione los productos que desea y finalmente confirme su pedido.")
	var selec string
	var pedido []Producto

	ciclo := true
	for ciclo {
		for k, v := range productos {
			fmt.Printf("%d. %s \t\tPrecio: $%0.2f\n", k, v.nombre, v.precio)
		}
		fmt.Printf("%d. CONFIRMAR PEDIDO\n", len(productos))
		fmt.Printf("%d. SALIR\n", (len(productos) + 1))

		fmt.Scanf("%s\n", &selec)
		prodIn, err := strconv.Atoi(strings.TrimSpace(selec))

		if err != nil {
			fmt.Println("ERROR: Debe ingresar un número de Producto válido.")
		} else {
			if prodIn == len(productos)+1 {
				fmt.Printf("Hasta pronto!...")
				ciclo = false
			} else if prodIn == len(productos) {
				if len(pedido) != 0 {
					return pedido
				} else {
					fmt.Println("Su pedido esta vacio!")
					break
				}
			} else if prodIn >= len(productos)+2 {
				fmt.Println("Opción invalida. Intente de nuevo")
			} else {
				fmt.Println("Producto seleccionado:", productos[prodIn].nombre)
				pedido = append(pedido, productos[prodIn])
			}
		}
	}
	return pedido
}

func FactibPedido(b []bool) bool {

	for _, data := range b {
		if data != true {
			return false
		}
	}
	return true
}

func ProcDispByProd(prodPedido Producto, ch chan<- Respuesta) (r Respuesta) {
	var resp Respuesta

	ingredientes.mutex.Lock()
	elDisp := make(map[int]int)

	// pasando slice a un map de ingredientes disponibles
	for _, s := range ingredientes.ingredientes {
		for _, d := range productos {
			if d.id == prodPedido.id {
				elDisp[s.id] = s.cantidad
			}
		}
	}
	elPedido := prodPedido.elementos

	for e := 0; e < len(elPedido); e++ {
		for i := 0; i < len(elDisp); i++ {
			// fmt.Println(e, "_", i, " - ", elPedido[e], "_", elDisp[i])
			if (e == i) && (elPedido[e] <= elDisp[i]) {
				resp.ok = 1
				resp.mensaje = "Se puede efectuar la preparacion del producto"
				resp.producto = prodPedido
				ch <- resp
			} else if (e == i) && (elPedido[e] > elDisp[i]) {
				resp.ok = 0
				resp.mensaje = "No hay suficiente cantidad de ingredientes para la preparacion del producto"
				resp.producto = prodPedido
				ch <- resp
			}
		}
	}

	// CONSULTAR PORQUÉ ESTA FORMA NO FUNCIONA CORRECTAMENTE.
	// for e, cant := range elPedido {
	// 	for i, cantIng := range elDisp {
	// 		fmt.Println(e, "_", i, " - ", cant, "_", cantIng)
	// 		if (e == i) && (cant <= cantIng) {
	// 			fmt.Printf(" -> ok\n")
	// 			resp.ok = 1
	// 			resp.mensaje = "Se puede efectuar la preparacion del producto"
	// 		} else if (e == i) && (cant < cantIng) {
	// 			resp.ok = 0
	// 			resp.mensaje = "No hay suficiente cantidad de ingredientes para la preparacion del producto"
	// 		}
	// 	}
	// }
	ch <- resp
	ingredientes.mutex.Unlock()
	return
}

func PrepairProduct(prodPedido Producto, chPrepair chan<- Respuesta, wgPrepair *sync.WaitGroup) {
	time.Sleep(time.Duration(prodPedido.tiempo) * time.Millisecond)
	var resp Respuesta

	elDisp := make(map[int]int)
	elPedido := prodPedido.elementos

	// pasando slice a un map de ingredientes disponibles
	for _, s := range ingredientes.ingredientes {
		for _, d := range productos {
			if d.id == prodPedido.id {
				elDisp[s.id] = s.cantidad
			}
		}
	}

	// Se verifica si la cantidad de ingredientes disponibles permiten la preparacion del produc. pedido
	for k, v := range elDisp {
		// Si no hay suficiente cantidad se genera respuesta de error.
		if v < elPedido[k] {
			resp.ok = -1
			resp.mensaje = "No hay suficiente cantidad de ingredientes para la preparacion del producto"
			resp.producto = prodPedido
			fmt.Println(resp)
			break
			// Panic
		}
	}

	// Si es posible preparar el producto
	if resp.ok != -1 {
		// Ciclos para actualizar las cantidades de ingredientes disponibles
		for i, ingreDisp := range ingredientes.ingredientes {
			for indElemDisp := range elDisp {
				if ingreDisp.id == indElemDisp {
					ingredientes.ingredientes[i].cantidad = elDisp[indElemDisp] - elPedido[indElemDisp]
				}
			}
		}

		// Se calcula el total a pagar.
		total = total + prodPedido.precio

		// Se genera respuesta ok=1 y se arroja mensaje de que el producto fue preparado
		resp.ok = 1
		resp.mensaje = "El Producto está listo!!!"
		resp.producto = prodPedido
		fmt.Println(resp)
	}

	// Se comunica al canal la respuesta sea que haya o todo salio bien.
	chPrepair <- resp
	return
}
