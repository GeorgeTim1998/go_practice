package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация больших чисел a и b
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("1048577", 10) // Устанавливает значение переменной a равным 1048577 в десятичной системе
	b.SetString("2097153", 10) // Устанавливает значение переменной b равным 2097153 в десятичной системе

	// Создание переменных для хранения результатов операций
	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)

	// Выполнение арифметических операций
	sum.Add(a, b) // Сложение
	sub.Sub(a, b) // Вычитание
	mul.Mul(a, b) // Умножение
	div.Div(a, b) // Деление

	// Вывод результатов
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", sub.String())
	fmt.Printf("a * b = %s\n", mul.String())
	fmt.Printf("a / b = %s\n", div.String())
}
