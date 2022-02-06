package main

import (
	"fmt"
)

var numero int
var numeroDos int

func main() {
	fmt.Println("Por favor digite un numero: ")
	fmt.Scanf("%d", &numero)

	fmt.Println("Por favor digite un numero: ")
	fmt.Scanf("%d", &numeroDos)

	fmt.Println("la suma es: ", numero+numeroDos)
}
