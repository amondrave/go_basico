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
			break // rompe la iteracion
		}
	}
	etiquetaEjemplo()
}

func etiquetaEjemplo() {
	i := 0
RUTINA: // Etiqueta de codigo
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a lanzar la rutina sumando de a 2")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}
