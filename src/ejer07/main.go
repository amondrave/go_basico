package main

import "fmt"

// las funciones tambien son tipo de datos
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(calculo(4, 3))

	// restamos , redifiniendo calculo

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Println(calculo(4, 3))

	// llamado a clousure

	tablaDelDos := 2
	miTabla := Tabla(tablaDelDos)
	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}

}

// Clousure
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
