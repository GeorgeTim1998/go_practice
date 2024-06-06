package main

import (
	"fmt"
	"strings"
)

func reverseWordOrder(input string) string {
	// разделяем слова по разделителю - пробелу
	words := strings.Split(input, " ")
	// буквально меняем слова местами идя с обоих концов массива к его середине
	// в середине все цикл остановится и перестановка будет завершена
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// возвращаем преобразованный в строку результат
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWordOrder(input)
	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reversed)
}
