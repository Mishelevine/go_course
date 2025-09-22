package main

import "fmt"

func task4() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	answer := 1
	devisions := 0

	for i := 2; i < num; i++ {
		devisions = 0
		if num%i != 0 {
			for j := 1; j <= i; j++ {
				if i%j == 0 && num%j == 0 {
					devisions += 1
				}
				if devisions > 1 {
					break
				}
			}
			if devisions == 1 {
				answer += 1
			}
		}
	}

	fmt.Printf("Функция Эйлера от числа %d: %d\n", num, answer)
}
