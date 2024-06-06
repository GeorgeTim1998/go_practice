package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Преобразуем строку в слайс рун
	runes := []rune(input)
	// буквально меняем их местами идя с обоих концов слайса к его середине
	// в середине все цикл остановится и перестановка будет завершена
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// возвращаем преобразованный в строку результат
	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reversed)
}
