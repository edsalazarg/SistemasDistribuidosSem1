package main

import "fmt"

func main() {
	// Calcular area triangulo
	var base float64
	var altura float64

	fmt.Println("Ingresa la base:")
	fmt.Scanf("%f\r", &base)

	fmt.Println("Ingresa la altura:")
	fmt.Scanf("%f\r", &altura)

	resultado := (base * altura) / 2

	fmt.Println("El area del triangulo es = ", resultado)
}
