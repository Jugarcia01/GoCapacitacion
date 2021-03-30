// 7. Dados los lados de un triángulo, decir qué tipo de triángulo es o si no es un triángulo válido.

package main

func main() {

}

func tipoTriangulo(lado1, lado2, lado3 float64) string {
	cantLados := 3
	var l1 = lado1
	var l2 = lado2
	var l3 = lado3

	if cantLados == 0 {
		return "La figura no tiene lados"
	} else {
		// return "La figura tiene uno o mas lados"
		if cantLados == 3 {
			// return "La figura es un triangulo"

			if (l1 > 0.0) && (l2 > 0.0) && (l3 > 0.0) {
				if l1 == l2 && l2 == l3 {
					return "El Triangulo es Equilatero"
				} else if l1 == l2 || l1 == l3 || l2 == l3 {
					return "El Triangulo es Isosceles"
				} else {
					return "El Triangulo es Escaleno"
				}
			} else {
				return "Los lados No tienen la longitud adecuada"
			}
		} else {
			return "La figura no corresponde a un triangulo"
		}
	}
}
