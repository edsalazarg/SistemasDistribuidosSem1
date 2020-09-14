package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.1416

func main() {
	// Calcular area triangulo
	var radio float64

	fmt.Println("Ingresa el radio:")
	fmt.Scanf("%f\r", &radio)

	// Sacando potencia de radio
	res1 := math.Pow(radio,2)

	resultado := pi * res1

	fmt.Println("El area del circulo es = ", math.Round(resultado*100)/100)
}
