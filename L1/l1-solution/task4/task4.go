package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

// Задача, которую будет исполнять каждая из горутин
func worker(channel chan int, wg *sync.WaitGroup) {
	// При завершении горутуны нужно оповестим об этом
	// wait группу
	defer func() {
		wg.Done()
	}()

	// Считываем данные из канал до тех пор,
	// пока канал channel открыт
	for data := range channel {
		fmt.Print(data, " ")
	}
}

// Создаёт пул из nWorkers горутин, который считывает
// данные из основоного потока и выводит в stdout,
// пока программа не получит прерывание. В случае
// прерывания, функция дожидается завршения всех горутин,
// чтобы избежать потери данных, отправленных в канал
func ReadData(nWorkers int) {
	// Создаём wait группу для nWorkers горутин
	var wg sync.WaitGroup
	wg.Add(nWorkers)
	channel := make(chan int)

	// Создаём nWorkers горутин и каждой
	// даём задачу по чтению данных из
	// канала channel
	for i := 0; i < nWorkers; i++ {
		go worker(channel, &wg)
	}

	// Создаём канал для фиксации сигналов
	sigChan := make(chan os.Signal, 1)
	// Объявляем, что будем фиксировать прерывание
	signal.Notify(sigChan, os.Interrupt)

	// Крутимся в цикле и кладём случайные значения в канал до тех пор,
	// пока не будет вызвано прерывание
	for {
		select {
		case channel <- rand.Intn(100):
		case <-sigChan:
			// В случае вызова прерывания
			// закрываем канал
			close(channel)
			// Ожидаем завершения всех работающих горутин
			wg.Wait()
			return
		}
	}
}

func main() {
	ReadData(250)
}
