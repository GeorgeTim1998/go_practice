package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker1(id int, done <-chan struct{}) {
	// рутина будет работать пока канал done не закрыт
	for {
		select {
		case <-done:
			fmt.Printf("Worker1 %d stopping\n", id)
			return
		default:
			fmt.Printf("Worker1 %d working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker2(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // получаем сообщение о таймауте
			fmt.Printf("Worker2 %d stopping\n", id)
			return
		default:
			fmt.Printf("Worker2 %d working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// часть третьего способа
var (
	stop = false
	mu   sync.Mutex
)

func worker3(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mu.Lock()
		if stop {
			mu.Unlock()
			fmt.Printf("Worker3 %d stopping\n", id)
			return
		}
		mu.Unlock()
		fmt.Printf("Worker3 %d working\n", id)
		time.Sleep(500 * time.Millisecond)
	}
}

func worker4(id int, wg *sync.WaitGroup, stoped <-chan struct{}) {
	// помечаем рутину как закончившую работу при помощи defer wg.Done(): уменьшаем счетчик активных го рутин
	defer wg.Done()

	// тело рутины выполняется пока не получится сигнал в stoped канале
	for {
		select {
		case <-stoped:
			fmt.Printf("Worker4 %d stopping\n", id)
			return
		default:
			fmt.Printf("Worker4 %d working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker5(id int, ch <-chan int) {
	// рутина работает пока канал не закрыт и выдает значения
	for val := range ch {
		fmt.Printf("Worker5 %d received: %d\n", id, val)
	}
	fmt.Printf("Worker5 %d stopping\n", id)
}

func main() {
	// #### первый способ
	// используем канал завершения го рутины done
	done := make(chan struct{})

	// запускаем рутины
	for i := 0; i < 3; i++ {
		go worker1(i, done)
	}

	time.Sleep(2 * time.Second)
	// закрываем канал
	close(done)
	time.Sleep(1 * time.Second) // Wait for goroutines to finish

	fmt.Println("\nExample 1 done\n")

	// #### второй способ
	// создаем го рутины с таймаутом. по его завершении выполнение прекратиться
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// создаем го рутины
	for i := 0; i < 3; i++ {
		go worker2(ctx, i)
	}

	time.Sleep(3 * time.Second) // Wait for goroutines to finish

	fmt.Println("\nExample 2 done\n")

	// #### третий способ
	// используем переменную стостояния stop, чтобы сообщить го рутине о том, что надо заершиться
	var wg1 sync.WaitGroup

	// создаем го рутины
	for i := 0; i < 3; i++ {
		wg1.Add(1)
		go worker3(i, &wg1)
	}

	time.Sleep(2 * time.Second)

	// записываем в переменную состояния информацию о том, что рутина должна остановиться
	mu.Lock()
	stop = true
	mu.Unlock()
	// ожидаем остановки го рутин
	wg1.Wait()

	fmt.Println("\nExample 3 done\n")

	// #### четвертый способ:
	// используем механизм sync.WaitGroup для ожидания остановки го рутины.
	// этот механизм позволяет основной программе дождаться выполнения всех го рутин
	var wg2 sync.WaitGroup
	stoped := make(chan struct{})

	// создаем го рутины и регистрируем их в sync.WaitGroup объекте
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go worker4(i, &wg2, stoped)
	}

	time.Sleep(2 * time.Second)
	close(stoped)

	// ожидаем выполнения всех го рутин
	wg2.Wait()

	fmt.Println("\nExample 4 done\n")

	// #### пятый способ:
	// прекрашение го рутины через закрываение канала, с которым го рутина работает
	ch := make(chan int)

	// начинаем работу 3х горутин
	for i := 0; i < 3; i++ {
		go worker5(i, ch)
	}

	// посылакем 5 значений в канал и закрываем его
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)

	time.Sleep(1 * time.Second) // ждем окончания работы го рутн

	fmt.Println("\nExample 5 done\n")

	fmt.Println("\nDemonstation done\n")
}
