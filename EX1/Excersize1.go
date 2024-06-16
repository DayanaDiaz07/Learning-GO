package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese lado 2: ")
	fmt.Scanln(&lado2)

	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(float64(lado1), 2) + math.Pow(float64(lado2), 2))
	perimetro := lado1 + lado2 + hipotenusa

	fmt.Printf("Área: %2f \n", area)
	fmt.Printf("Perímetro: %2f", perimetro)
}
