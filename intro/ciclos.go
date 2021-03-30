package main

import "fmt"

func main() {

	fmt.Println("Ciclo for con inicializac, condicional e incremento")
	cont := 10
	for i := 0; i <= cont; i++ {
		fmt.Println(i)
	}

	fmt.Println("Ciclo for con condicional inicializac externa e incremento interno")
	ii := 0
	for ii <= cont {
		fmt.Println(ii)
		ii++
	}

	fmt.Println("Ciclo for con condicional e incremento internos, inicializac externa")
	iii := 0
	for {
		fmt.Println(iii)
		iii++
		if iii >= 10 {
			break
		}
	}

	fmt.Println("Ciclo con continue y break")
	i2 := 0
	for {
		if i2 == 10 {
			i2++ // Si no se fija aquí entonces se quedaba en ciclo infinito
			continue
		}
		fmt.Println(i2)
		i2++ // Si el incremento no estuviera aquí entonces quedaría en un ciclo infinito
		if i2 >= 20 {
			break
		}
	}

}
