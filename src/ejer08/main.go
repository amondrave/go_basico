package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var arreglo []int

func main() {
	llenarArreglo(10)
	fmt.Println(arreglo)
	arregloConImpares := arregloSoloImpares(arreglo)
	fmt.Println(arregloConImpares)
}

func llenarArreglo(tamanio int) {
	arreglo = make([]int, tamanio)
	for i := 0; i < tamanio; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(20))
		arreglo[i] = int(result.Int64())
	}
}

func arregloSoloImpares(arreglo []int) []int {
	arregloRetorno := make([]int, 0)
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i]%2 != 0 {
			arregloRetorno = append(arregloRetorno, arreglo[i]) // append tiene costo en proceso
		}
	}
	return arregloRetorno
}
