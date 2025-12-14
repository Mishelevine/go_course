package main

import (
	"context"
	"fmt"
	"time"
)

func check(ctx context.Context, n int, res chan<- bool) {
	for n > 1 {
		// Проверка ограничения времени
		select {
		case <-ctx.Done():
			res <- false
			return
		default:
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	// Проверили все
	res <- true
}

func worker(ctx context.Context, numbers <-chan int, checked chan<- int) {
	count := 0
	running := true

	for running {
		select {
		// Проверка ограничения времени
		case <-ctx.Done():
			running = false
		case n, ok := <-numbers:
			// Проверка что канал с числами не закрыт
			if !ok {
				running = false
				break
			}

			// Проверка числа
			res := make(chan bool, 1)
			go check(ctx, n, res)

			// Ожидание результата проверки или завершения таймера
			select {
			case <-ctx.Done():
				running = false
			case <-res:
				count++
			}
		}
	}

	// Возврат кол-ва проверенных
	checked <- count
}

func task2_5() {
	const limit = 3000000      // можно экспериментировать с верхней границей
	const workersCnt = 8       // можно экспериментировать с количеством горутин
	nums := make(chan int, 10) // буфер размера 10

	// Горутина-генератор чисел
	go func() {
		for i := 1; i <= limit; i++ {
			nums <- i
		}
		close(nums)
	}()

	checked := make(chan int, workersCnt)

	// Ограничение по общему времени работы программы
	ctxTO, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// Запуск воркеров
	for k := 0; k < workersCnt; k++ {
		go worker(ctxTO, nums, checked)
	}

	// Суммируем, сколько каждый воркер успел проверить
	cnt := 0
	for j := 0; j < workersCnt; j++ {
		cnt += <-checked
	}

	fmt.Printf("\nВсего проверено %d чисел\n", cnt)
}
