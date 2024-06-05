package main

// проблема исходного кода
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }
//
// заключается в том, что операция v[:100] создает строку, которая указывает на оригинал
// это означает, что если нам надо работать только с первыми 100 символами большой строки, то
// большая строка будет существовать тоже в памяти. Это неэффективно
// чтобы этого избежать, надо создать копию нужной части строки. GC затем удалит исходную строку за ненадобностью
// и использование памяти будет более оптимальным.

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	// создается большая строка
	v := createHugeString(1 << 30)

	// создаем копию нужных символов исходной строки через копирование copy() в массив byte
	neededString := make([]byte, 100)
	copy(neededString, v[:100])

	// преобразуем массив в строку. Теперь мы работает только с 100 символами оригинальной строки
	justString = string(neededString)

	fmt.Println(justString)
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func main() {
	someFunc()
}
