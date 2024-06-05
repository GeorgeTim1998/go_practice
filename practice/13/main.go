package main

import "fmt"

func main() {
	a, b := 5, 10

	// Значения переменных до свопа
	fmt.Println("Before swap: a =", a, ", b =", b)

	// Обмен значениями переменных без использования временной переменной
	a, b = b, a

	// Значения переменных после свопа
	fmt.Println("After swap: a =", a, ", b =", b)
}
