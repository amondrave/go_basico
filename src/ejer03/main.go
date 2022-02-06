package main

import "fmt"

var estado bool

func main() {
	estado = true
	if estado {
		fmt.Println("Hola esto es verdadero")
	} else {
		fmt.Println("esto es falso")
	}
	fmt.Println("Imprime igual")
}
