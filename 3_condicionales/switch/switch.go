package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("Escribe ", i, " como ")
	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Es día entre semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Es antes del medio día")
	default:
		fmt.Println("Ya pasa del medio día")
	}

	queSoy := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Soy un booleano")
		case int:
			fmt.Println("Soy un entero")
		default:
			fmt.Printf("No sé el tipo de dato que soy %T\n", t)
		}
	}
	queSoy(true)
	queSoy(1)
	queSoy("Hey!")
}
