// 6. Calcular qué hora será dentro de N horas. Ej: Si son las 16, dentro de N=12, serán las 4;
//    dentro de N=24, serán las 16; dentro de N=28, serán las 20.

package main

import (
	"fmt"
	"time"
)

func main() {
	var horasAdic = 4

	calcDifHoras(horasAdic)

}

func calcDifHoras(cantHorasFut int) {
	loc, _ := time.LoadLocation("America/Bogota")
	fechHoraActual := time.Now().In(loc)
	fmt.Println("\nUbicación: ", loc, " Fecha y Hora :\n", fechHoraActual)

	horasAdic := int(time.Now().Hour()) + cantHorasFut

	fechaFut := time.Date(2021, time.March, 29, horasAdic, 00, 00, 111, loc)
	fmt.Println("Ubicación:", loc, " Fecha y Hora Futura Aproximada: \n", fechaFut)
	diff := fechaFut.Sub(fechHoraActual)

	hrs := float32(diff.Hours())
	fmt.Printf("Dif. en Horas : %0.2f Hours\n", hrs)

	dias := int(diff.Hours() / 24)
	fmt.Printf("Dif. en Dias : %d days\n", dias)

}
