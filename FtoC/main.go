package main

import (
	"fmt"
)

func main() {
	// Calcular area triangulo
	var grados float64

	fmt.Println("Ingresa el los grados en Fahrenheit:")
	fmt.Scanf("%f\r", &grados)

	resultado := (grados-32) * 5/9

	fmt.Println("Los grados en Centigrados son = ", resultado)
}
