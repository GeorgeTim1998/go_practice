package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция для воркера
// Используем операторы <-chan чтобы показать, что каналы в воркере будут использоваться только для получения данных.
func worker(id int, data <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num := <-data:
			fmt.Printf("Worker %d received: %d\n", id, num)
		case <-done:
			fmt.Printf("Worker %d exiting\n", id)
			return
		}
	}
}

func main() {
	// Читаем количество воркеров из аргументов командной строки
	var numWorkers int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&numWorkers)

	// Создаем канал для посылки данных в него,
	// а также канал, куда пошлем сообщение о окончании работы программы.
	// Также создаем объект синхронизации
	data := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Запускаем воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, data, done, &wg)
	}

	// Запуск горутины для генерации данных
	go func() {
		for {
			select {
			case <-done:
				// Когда программа завершается - закрываем канал для получения данных
				close(data)
				return
			default:
				data <- rand.Intn(100)             // Посылаем рандомные данные в канал
				time.Sleep(100 * time.Millisecond) // Имитация задержки
			}
		}
	}()

	// Захват сигнала завершения программы (Ctrl+C).
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Блокируем дальнейшее выполнение main ожиданием сигнала окончания программы.
	<-sigChan

	close(done) // Отправляем сигнал завершения

	// Ждем завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers exited. Exiting main.")
}

// Обоснование реализованного способа окончания программы:
// Используется канал done для окончания выполнения программы, чтобы корректно завершить выполнение программы
// и всех ее воркеров.
// Канал done используется для гарантированного завершения цикла выполнения в воркере и завершения его работы
