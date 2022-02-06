package main

import (
	"fmt"
	"strconv"
)

var numero int   // el scope es de todo el paquete
var texto string // Si la primera letra es minuscula la variable es solo local
var status bool
var cadena string

var moneda int64 = 9

func main() {
	numeroEntero := 4 // Scope solo funcion : solo es para la primera asignacion
	var numeroSuma, numeroSumaDos int
	numeroSuma, numeroSumaDos = 5, 6
	fmt.Println(numeroSuma + numeroSumaDos)
	fmt.Println(numeroEntero)
	numero = int(moneda)               // convertir entero 64 a entero general
	texto = fmt.Sprintf("%d", moneda)  // Entero a String
	cadena = strconv.Itoa(int(moneda)) // Otra manera de convertir a string con un paquete
	mostrarStatus()
	mostrarVariables()
}

func mostrarStatus() {
	fmt.Println(status)
}

func mostrarVariables() {
	fmt.Println(numero)
	fmt.Println(texto)
	fmt.Println(cadena)
}
