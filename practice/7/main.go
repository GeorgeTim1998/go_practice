package main

import (
	"fmt"
	"sync"
)

func main() {
	// создаем объект для ожидания завершения го рутин
	var wg sync.WaitGroup

	// создаем объект map
	m := make(map[int]int)

	// создаем mutex объект
	var mu sync.Mutex

	// Функция для записи данных в map.
	// она во время записи блокирует объект map для изменений другими горутинами через mutex
	writeToMap := func(key, value int) {
		defer wg.Done()
		mu.Lock()
		m[key] = value
		mu.Unlock()
	}

	// Запуск нескольких горутин для конкурентной записи в map с регистрацией в объекте sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(i, i*i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод содержимого map
	for k, v := range m {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}
}
