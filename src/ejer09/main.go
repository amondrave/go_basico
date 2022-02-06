package main

import "fmt"

// Mapa es una estructura de elementos que se pueden serializar
// se rigen por clave valor

func main() {
	mapa := crearMapaCapitales()
	fmt.Println(mapa)
	fmt.Println(mapa["COLOMBIA"])

	mapa2 := crearMapaCampeonato()

	fmt.Println(mapa2)

	recorrerMapa(mapa2)
}

// Creando mapa copn make
func crearMapaCapitales() map[string]string {
	paises := make(map[string]string)
	paises["MEXICO"] = "Distrito Capital"
	paises["COLOMBIA"] = "Bogota"
	return paises
}

// Crear mapa directamente sin make
func crearMapaCampeonato() map[string]int {
	campeonato := map[string]int{
		"Real Madrid": 30,
		"Barcelona":   24,
		"Valencia":    15}
	campeonato["Villarreal"] = 17 // puedo agregar de manera natural al mapa

	return campeonato
}

func recorrerMapa(mapa map[string]int) {
	for equipo, puntaje := range mapa {
		fmt.Printf("Equipo %s , puntaje %d \n", equipo, puntaje)
	}
}
