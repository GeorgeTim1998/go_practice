package main

import (
	"fmt"
)

func main() {
	// Исходная последовательность температур
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Карта для группировки температур
	tempGroups := make(map[int][]float64)

	// Группируем температуры с шагом 10 градусов
	for _, temp := range temperatures {
		// Вычисляем ключ для текущей температуры посредством окрушления вниз до ближайщего целого десятка
		key := int(temp/10) * 10

		// Добавляем температуру в соответствующую группу добавлением в массив, который отвечает ключу-диапазону в map
		tempGroups[key] = append(tempGroups[key], temp)
	}

	// Выводим группы температур
	for key, temps := range tempGroups {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
