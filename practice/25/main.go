package main

import (
	"fmt"
	"time"
)

// MySleep приостанавливает выполнение программы на заданное количество времени.
func MySleep(d time.Duration) {
	// time.After(d) возвращает канал, в который придет значение только через d секунд
	// и это ожидание значение блокирует выполнение программы
	<-time.After(d)
}

func main() {
	fmt.Println("Start sleeping...")
	MySleep(10 * time.Second)
	fmt.Println("Woke up after 10 seconds")
}
