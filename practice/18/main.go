package main

import (
	"fmt"
	"sync"
)

// Counter - структура-счетчик
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment - метод для увеличения счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value - метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// методы Increment и Value оба используют Mutex, чтобы защититься от гонки состояний

func main() {
	// Добавляем объект синхронизации
	var wg sync.WaitGroup
	counter := Counter{}

	// выбираем число го рутин и количество увеличений, которое должно произойти в каждой рутине
	numGoroutines := 100
	incrementsPerGoroutine := 1000

	// Запуск конкурентных горутин
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				// увеличиваем счетчик
				counter.Increment()
			}
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод итогового значения счетчика
	fmt.Println("Final counter value:", counter.Value())
}
