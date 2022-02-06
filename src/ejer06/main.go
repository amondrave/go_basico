package main

import "fmt"

func main() {
	fmt.Println(numeroMultiplicadoPorDos(5))
	fmt.Println(saludo("Angel"))
	var resultado, esPar = retornarDosElementos(10)
	fmt.Printf("El resultado: %d es par? %t \n", resultado, esPar)
	calculo(5, 6, 7, 8)
}

func numeroMultiplicadoPorDos(numero int) int {
	return numero * 2
}

func saludo(nombre string) string {
	return "Hola " + nombre + " Como estas ?"
}

func retornarDosElementos(numero int) (int, bool) {
	if numero%2 == 0 {
		return numero * 2, true
	} else {
		return numero, false
	}
}

/**
* GO no tiene sobrecarga de metodos, hay parametros dinamicos
 */

func calculo(numero ...int) {
	total := 0
	for _, num := range numero { // range instruccion devuelve dos valores, indice del elemento y el valor de este
		total += num
	}
	println("el valor total es: ", total)
}
