package main

import "fmt"

func removeElement(slice []int, i int) []int {
	// Проверяем, что индекс i находится в пределах слайса
	if i < 0 || i >= len(slice) {
		return slice
	}
	// Объединяем две части слайса: до i-го элемента и после i-го элемента
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice)

	// Удаляем элемент с индексом 3 (значение 4)
	indexToRemove := 3
	slice = removeElement(slice, indexToRemove)

	fmt.Println("Modified slice:", slice)
}
