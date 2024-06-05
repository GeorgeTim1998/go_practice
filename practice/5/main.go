package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(ch chan<- int, done <-chan struct{}, received <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		select {
		case ch <- i:
			// посылаем сообщение в канал и ожидаем его получения
			fmt.Println("\nSent:", i)
			fmt.Println("Await receive by the receiver")
			<-received
		case <-done: // ловим сообщение о том, что программа завершает свою работу по таймауту
			// закрываем канал для посылки значений и завершаем го рутину
			fmt.Println("\nClose the channel for values")
			close(ch)
			return
		}
	}
}

func receiver(ch <-chan int, done <-chan struct{}, received chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				// рутина sender завершила свою работу. Теперь завершаем работу receiver
				fmt.Println("Channel already closed. Closing 'received' channel.")
				// закрываем канал сообщения между sender и receiver и завершаем выполнение го рутины
				close(received)
				return
			}

			fmt.Println("\nReceived:", val)
			received <- struct{}{}
		}
	}
}

func main() {
	// слздаем канал для посылки значений
	ch := make(chan int, 1)

	// создаем канал, через который receiver и sender будут общаться. Sender не будет посылать сообщения пока Receiver его не получил
	received := make(chan struct{}, 1)

	// добавляем обхект синхронизации чтобы дождаться выполнения/завершения рутин
	var wg sync.WaitGroup

	// Канал, через который будет сообщаться о конце работы программы
	done := make(chan struct{})

	// ставим таймаут для выполнения программы
	time.AfterFunc(1*time.Second, func() {
		close(done)
	})

	// запускаем го рутины
	wg.Add(2) // Увеличиваем счетчик ожидания на 2, так как у нас две горутины
	go sender(ch, done, received, &wg)
	go receiver(ch, done, received, &wg)

	wg.Wait() // Ожидание завершения работы обеих горутин
	fmt.Println("\nProgram completed")
}
