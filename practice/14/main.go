package main

import (
	"fmt"
)

func main() {
	// Примеры переменных разных типов
	var a int = 42
	var b string = "hello"
	var c bool = true
	var d chan int = make(chan int)

	// Определяем типы переменных
	determineType(a)
	determineType(b)
	determineType(c)
	determineType(d)
}

// Функция для определения типа переменной
func determineType(v interface{}) {
	// v.(type) - Type assertion - достаем базовый тип у интерфейса
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("unknown")
	}
}
