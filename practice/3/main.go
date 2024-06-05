package main

import (
	"fmt"
	"sync"
)

func main() {
	integers := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	// Добавляем объект синхронизации
	var wg sync.WaitGroup

	// Создаем го рутину для того чтобы асинхронно получать результаты вычисления квадрата числа и суммировать их
	// по мере поступления.
	// Для этого итерируемся по всем значениям, которые будут попадать в канал
	go func() {
		sum := 0
		for result := range c {
			sum += result
		}
		fmt.Println("Sum of squares:", sum)
	}()

	// запускаем все го рутины вместе с объектом синхронизации, чтобы потом понять когда канал можно будет закрыть
	for _, i := range integers {
		wg.Add(1)
		go square(i, c, &wg)
	}

	// Закрываем канал только после завершения всех горутин
	go func() {
		wg.Wait()
		close(c)
	}()

	// Заставляем main рутину ждать выполнения остальных
	wg.Wait()
}

// Считаем квадрат и отмечаем в объекте синхронизации что го рутина закончила свою работу
func square(a int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- a * a
}
