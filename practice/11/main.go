package main

import (
	"fmt"
)

func main() {
	// Два исходных множества
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Получаем пересечение множеств
	intersection := intersect(set1, set2)

	// Выводим результат
	fmt.Println("Intersection:", intersection)
}

// Функция для нахождения пересечения двух множеств
func intersect(set1, set2 []int) []int {
	// Создаем карты для представления множеств
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	// Заполняем первую карту элементами первого множества
	for _, v := range set1 {
		map1[v] = true
	}

	// Заполняем вторую карту элементами второго множества
	for _, v := range set2 {
		map2[v] = true
	}

	// Находим пересечение множеств
	var intersection []int
	for k := range map1 {
		// если эдемент первого множества есть во втором - добавляем в их пересечение
		if map2[k] {
			intersection = append(intersection, k)
		}
	}

	return intersection
}
