package main

import (
	"fmt"
)

func task2_2() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	ch := make(chan int, 4)
	defer close(ch)

	numWindow := n / 4

	go func() {
		var sum1 int
		for i := 0; i < numWindow; i++ {
			sum1 += numbers[i]
		}
		ch <- sum1
	}()

	go func() {
		var sum2 int
		for i := numWindow; i < numWindow*2; i++ {
			sum2 += numbers[i]
		}
		ch <- sum2
	}()

	go func() {
		var sum3 int
		for i := numWindow * 2; i < numWindow*3; i++ {
			sum3 += numbers[i]
		}
		ch <- sum3
	}()

	go func() {
		var sum4 int
		for i := numWindow * 3; i < n; i++ {
			sum4 += numbers[i]
		}
		ch <- sum4
	}()

	var result int
	for i := 0; i < 4; i++ {
		result += <-ch
	}

	fmt.Printf("Сумма чисел от 0 до %v = %v\n", n, result)
}
