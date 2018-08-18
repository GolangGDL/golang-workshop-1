package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func substract(a, b int) (int, error) {
	var err error
	if a < b {
		err = errors.New("error: Primer valor mayor que el segundo")
		return 0, err
	} else {
		err = nil
	}
	result := a - b
	return result, err
}

func main() {
	fmt.Println(" Calculadora ")
	resAdd := add(2, 4)
	resSub, err := substract(10, resAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("2+4=", resAdd)
	fmt.Println("10-(2+4)=", resSub)
	resSub, err = substract(10, 20)
	if err != nil {
		fmt.Println(err)
	}
}
