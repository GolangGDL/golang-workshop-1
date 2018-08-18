package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func main() {
	fmt.Println(persona{"Pepe", 20})
	fmt.Println(persona{nombre: "Toño", edad: 30})
	fmt.Println(persona{nombre: "María"})
	fmt.Println(&persona{nombre: "Pablo", edad: 40})
	s := persona{nombre: "José", edad: 50}
	fmt.Println(s.nombre)
	sp := &s
	fmt.Println(sp.edad)
	sp.edad = 51
	fmt.Println(sp.edad)
}
