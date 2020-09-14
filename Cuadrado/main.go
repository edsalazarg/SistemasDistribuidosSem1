package main

import "fmt"

func main() {
	// Calcular area cuadrado
	var base float64

	fmt.Println("Ingresa el tamaño de un lado del cuadrado:")
	fmt.Scanf("%f", &base)

	resultado := base * base

	fmt.Println("El area del cuadrado es = ", resultado)
}
