package main

import (
	"fmt"
	"sync"
)

type Pair struct {
	value int
	index int
}

// Выводи в stdout квадраты элементов среза.
// Порядок вывода чисел не сохраняется
func PrintScuares(arr []int) {
	var wg sync.WaitGroup

	// Проходим по каждому элемента среза
	for _, number := range arr {

		// Увеличиваем количество горутин в wait группе,
		// завершения которых нужно дождаться
		wg.Add(1)

		// Создаём горутину и запускаем в ней задачу, используя анонимную функцию
		go func(numb int) {
			// Создаём отложенный вызов с оповещением wait группы
			// о завершении работы горутины
			defer wg.Done()

			// Выводим ответ в stdout
			fmt.Print(numb*numb, " ")
		}(number)
	}

	// Дожидаемся
	wg.Wait()
	fmt.Println()
}

func PrintOrderedSquares(arr []int) {
	var wg sync.WaitGroup
	squaredVals := make([]int, len(arr))

	channel := make(chan Pair)

	for idx, number := range arr {

		// Увеличиваем количество горутин в wait группе,
		// завершения которых нужно дождаться
		wg.Add(1)

		// Создаём горутину и запускаем в ней задачу, используя анонимную функцию
		go func(numb int, ch chan Pair, i int) {
			// Создаём отложенный вызов с оповещением wait группы
			// о завершении работы горутины
			defer wg.Done()

			// Передаём пару значение - индекс в канал
			ch <- Pair{
				value: numb * numb,
				index: i,
			}
		}(number, channel, idx)
	}

	for i := 0; i < len(arr); i++ {
		// Достаём пару значение - индекс из канала
		data := <-channel
		squaredVals[data.index] = data.value
	}

	// Выводим полученные значения в нужном порядке
	for _, item := range squaredVals {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	// Выводит, возведенные в квадрат, значения без сохранения порядка
	PrintScuares(arr)

	// Выводит, возведенные в квадрат значения с сохранением порядка
	PrintOrderedSquares(arr)
}
