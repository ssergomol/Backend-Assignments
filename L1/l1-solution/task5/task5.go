package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Последовтельно отправляет значения в канал.
// С другого конца канала происходит чтение и вывод
// в stdout. По истечению времени sec программа завершается.
func ReadUntilTimeExceed(sec int) {
	var wg sync.WaitGroup
	wg.Add(1)
	channel := make(chan int)
	done := make(chan bool)

	// Вызываем в функцию, которая отправляет текущее время
	// в канал timeout через sec секунд
	timeout := time.After(time.Duration(sec) * time.Second)

	go func() {
		// При завершении горутины оповестим об этом
		// основной поток через канал done
		defer func() {
			done <- true
		}()

		// Достаём данные из канала до тех пор,
		// пока не получено значение из канала timeout
		for {
			select {
			case <-timeout:
				return
			case data := <-channel:
				fmt.Print(data, " ")
			}
		}
	}()

	// Пока горутина не завершилась, кладём случайные
	// значения в канал.
	for {
		select {
		case <-done:
			return

		case channel <- rand.Intn(100):
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func main() {
	ReadUntilTimeExceed(2)
}
