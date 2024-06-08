package main

import (
	"fmt"
)

// Quick sort (быстрая сортировка) – суть алгоритма заключается в разделении массива на два под-массива,
// средней линией считается элемент, который находится в самом центре массива.
// В ходе работы алгоритма элементы, меньшие чем средний будут перемещены в лево, а большие в право.
// Такое же действие будет происходить рекурсивно и с под-массива, они будут разделяться на еще два под-массива до тех пор,
// пока не будет чего разделать (останется один элемент).

func quicksort(arr []int) []int {
	fmt.Printf("\n\n")
	fmt.Printf("%2d ", arr)
	fmt.Printf("\n")
	// если нечего сортировать (в массиве один и меньше элементов) - возвращаем исходный массив
	if len(arr) < 2 {
		return arr
	}

	// в ходе выполнения этой функции мы хотим частично отсортировать переданный массив следующим образом:
	// мы хотим выбрать средний элемент массива (pivot) и расположить элементы массива относительно этого pivot элемента так,
	// чтобы слева от pivot оказались только элементы меньшие pivot, а справа от pivot - элементы большие.
	// Проблема заключается в том, что мы не знаем, где должен находиться элемент pivot в массиве,
	// чтобы удовлетворить тому состоянию массива, которое мы хотим получить после частичной сортировки.
	// в итоге, за один прогон массива мы хотим решить две задачи:
	// - найти где должен находиться pivot, чтобы удовлетворить частично отсортированному массиву после выполнения этой функции
	// - переставить все остальные элементы относительно pivot элемента, чтобы слева от него были элементы меньшие, а справа - большие.

	// индекс положения pivot элемента в массиве обозначим как left. его истинное значение предстоит определить в ходе выполнения данной функции
	// в самом начале алгоритма, предположим, что индекс left равен 0 (pivot должен находиться в начале массива)
	// в переменную rigth всегда будем использовать как обозначение индекса последнего элемента в переданном массиве arr
	left, right := 0, len(arr)-1

	// выбирем pivot как средний элемент массива
	pivotIndex := len(arr) / 2

	// по скольку мы не знаем, где должен находиться находиться pivot,
	// чтобы удовлетворить нашему желаемому состоянию массива после частичной сортировки, поместим его в конец массива
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	fmt.Printf("%2d ", arr)
	fmt.Printf("\n")
	// тут мы проходим по всем элементам массива arr
	for i := range arr {
		// в этом блоке if мы проверяем как соотносится текущий элемент массива arr[i] и значение pivot элемента arr[right]
		if arr[i] < arr[right] {
			// если элемент arr[i] меньше чем pivot элемент, то это означает, что этот элемент должен идти точно до pivot элемента в массиве после нашей частичной сортировки
			// мы ставим меньший чем pivot элемент на место предполагаемой изначально позиции pivot элемента, которая нам еще не известа и обозначена как left и имеет текущее значение - 0
			arr[i], arr[left] = arr[left], arr[i]
			fmt.Printf("%2d ", arr)
			fmt.Printf("\n")

			// т.к мы нашли элемент меньший чем pivot, то изначально предполагаемая позиция left = 0 - не верна (там теперь стоит найденный меньший элемент),
			// и значение позиции надо увеличить на 1
			// при каждом новом нахождении элемента меньшего чем pivot, истинная позиция left будет смещаться на 1, что в конце прогона по массиву гарантирует,
			// что все меньшие элементы чем pivot будут слева от позиции left в массиве, а большее элементы справа от этой позиции
			// само значение left будет отвечать позиции pivot элемента, которую мы и хотели определить
			left++
		}
	}

	// Возвращаем пивот на его место (найденное место left), которое удовлетворит нашему желанию получить массив, в котором слева от pivot элемента все элементы массива меньше его,
	// а справа - больше
	arr[left], arr[right] = arr[right], arr[left]
	fmt.Printf("%2d ", arr)

	// Рекурсивно сортируем левую и правую части аналогичным образом
	// левая часть это до позиции left не включая left - там находятся все элементы меньшие, чем pivot, выбранный на текущей итерации
	quicksort(arr[:left])
	// правая часть это от позиции left + 1 и до конца массива arr - там находятся все элементы большие, чем pivot, выбранный на текущей итерации
	quicksort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{9, 10, 3, 2, 4, 6, 5, 7}
	fmt.Println("Original array:", arr)

	sortedArr := quicksort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
