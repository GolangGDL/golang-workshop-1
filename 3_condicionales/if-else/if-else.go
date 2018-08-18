package main

import "fmt"

func main() {

	if 8%4 == 0 {
		fmt.Println("8 es divisible entre 4")
	}

	a := 10
	if a > 5 {
		fmt.Println("a es mayor a 5")
	} else {
		fmt.Println("a es menor o igual a 5")
	}

	if num := 10; num < 0 {
		fmt.Println(num, "es negativo")
	} else if num < 10 {
		fmt.Println(num, "tiene 1 dígito")
	} else {
		fmt.Println(num, "tiene múltiples dígitos")
	}
}
