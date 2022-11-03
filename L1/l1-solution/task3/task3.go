package main

import "fmt"

// Конкурентно считает сумму квадратов элементов среза
func SumOfSquares(arr []int) int {
	channel := make(chan int)
	var sum int

	for _, number := range arr {

		// Создаём горутину и запускаем в ней задачу, используя анонимную функцию
		go func(numb int, ch chan int) {

			// Кладём квадрат элемента среза в канал
			ch <- numb * numb
		}(number, channel)
	}

	for i := 0; i < len(arr); i++ {
		// Достаём значение из канала и добавляем к сумме
		data := <-channel
		sum += data
	}

	return sum
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(SumOfSquares(arr))
}
