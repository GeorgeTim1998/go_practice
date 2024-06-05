package main

import (
	"fmt"
)

// setBit устанавливает i-й бит в переменной num в значение value (1 или 0).
func setBit(num int64, i int, value int) int64 {
	if value == 1 {
		// Устанавливаем i-й бит в 1
		num |= (1 << i) // присвоение побитового 'или' с числом (1 << i)
	} else if value == 0 {
		// Устанавливаем i-й бит в 0
		num &= ^(1 << i) // пирсвоение побитового 'и' с числом ^(1 << i)
	} else {
		fmt.Println("Value must be 0 or 1")
	}
	return num
}

func main() {
	var num int64 = 64 // Пример переменной
	fmt.Printf("%30s: %064b\n", "Initial value", num)

	// Устанавливаем 1-й бит в 1
	num = setBit(num, 1, 1)
	fmt.Printf("%30s: %064b\n", "After setting bit 1 to 1", num)

	// Устанавливаем 1-й бит в 0
	num = setBit(num, 1, 0)
	fmt.Printf("%30s: %064b\n", "After setting bit 1 to 0", num)
}
