package main

import (
	"errors"
)

type Cuenta struct {
	Id    string
	Monto float32
}

func main() {

}

func (c *Cuenta) Consignar(valor float32) (monto float32, err error) {
	if valor < 0.0 {
		err = errors.New("El valor no debe ser negativo")
		return
	}
	c.Monto += valor
	monto = c.Monto
	return
}

// func (c *Cuenta) Retirar(valor float32) (monto float32, err error) {
// 	if c.Monto < valor {
// 		err = errors.New("El valor excede el monto actual de la cuenta")
// 		return
// 	}
// 	if valor < 0 {
// 		err = errors.New("El valor no debe ser negativo")
// 		return
// 	}
// 	c.Monto -= valor
// 	monto = c.Monto
// 	return
// }
