package main

import (
	"fmt"
)

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем пустое множество
	set := make(map[string]struct{})

	// Добавляем элементы в множество
	// создание собственного множества достигается за счет того, что ключи в map уникальны и повторы автоматически уберутся
	// эффективности такому способу добавляет тот факт, что struct{}{} не занимает памяти
	for _, str := range strings {
		set[str] = struct{}{}
	}

	// Выводим элементы множества
	for key := range set {
		fmt.Println(key)
	}
}
