package main

import (
	"fmt"
)

func calcNumSum(num int) int {
	var sum int

	numStr := fmt.Sprintf("%d", num)
	length := len(numStr)

	for i := 0; i < length; i++ {
		digit := int(numStr[i]) - '0'

		sum += digit
	}

	return sum
}

func task2() {
	var num int
	fmt.Print("Введите неотрицательное число: ")
	fmt.Scan(&num)
	fmt.Printf("Сумма элементов введенного числа: %d\n", calcNumSum(num))
}
