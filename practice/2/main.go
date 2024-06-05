package main

import (
	"fmt"
	"sync"
)

func main() {
	// Задачем массив с числами
	integers := [5]int{2, 4, 6, 8, 10}

	// Добавляем объект синхронизации
	var wg sync.WaitGroup

	// В цикле создаем го рутины и увеличиваем счетчик го рутин у wg
	for _, i := range integers {
		wg.Add(1)
		go square(i, &wg)
	}

	// Ожидаем выполнения всех го рутин
	wg.Wait()
}

func square(a int, wg *sync.WaitGroup) {
	// Уменьшаем счетчик активных го рутин
	defer wg.Done()

	fmt.Println(a * a)
}
