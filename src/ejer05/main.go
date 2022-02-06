package main

import "fmt"

func main() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("El valor de i es: ", i)

	}
	numero := 0
	for {
		fmt.Println("adivine el numero para romper el ciclo: ")
		fmt.Scanf("%d", &numero)
		if numero == 5 {
			break
		}
	}
}
