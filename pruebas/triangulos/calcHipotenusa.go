// 3. Calcular la hipotenusa dados los catetos
// h = raizCuadrada( b2 + c2)

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	cateto1 := 3.00
	cateto2 := 4.00
	var hipotenusa float64

	//	math.Sqrt( math.Pow(cat1) + math.Pow(cat2))
	hipotenusa, _ = calcHipotenusa(cateto1, cateto2)

	fmt.Println(hipotenusa)

}

func calcHipotenusa(cat1, cat2 float64) (hipotenusa float64, err error) {

	if cat1 <= 0 || cat2 <= 0 {
		err = errors.New("No es válido calcular hipotenusa con valores de long. de catetos 0 o negativos")
	} else {
		hipotenusa = math.Sqrt(math.Pow(cat1, 2.0) + math.Pow(cat2, 2.0))
	}
	return
}
