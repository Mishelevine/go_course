package main

import (
	"fmt"
)

const (
	n = 2
	m = 3
)

func reduction(a [n][m]float64) [n][m]float64 {
	// Редукция по строкам
	for i := 0; i < n; i++ {
		// Ищем мин
		min := a[i][0]
		for j := 0; j < m; j++ {
			if min > a[i][j] {
				min = a[i][j]
			}
		}
		// Редукция
		for j := 0; j < m; j++ {
			a[i][j] -= min
		}
	}

	// Редукция по столбцам
	for i := 0; i < m; i++ {
		// Ищем мин
		min := a[0][i]
		for j := 0; j < n; j++ {
			if min > a[j][i] {
				min = a[j][i]
			}
		}
		// Редукция
		for j := 0; j < n; j++ {
			a[j][i] -= min
		}
	}

	return a
}

func task6() {
	fmt.Printf("Введите чисела для заполнения матрицы %v на %v\n", n, m)
	var matrix [n][m]float64

	for i := 1; i <= n; i++ {
		fmt.Printf("Введите %v чисел для строки %v (числа вводите через пробел)", m, i)

		for j := range m {
			fmt.Scan(&matrix[i-1][j])
		}
	}

	fmt.Println("Изначальная матрица:")
	fmt.Printf("%v\n", matrix)

	result := reduction(matrix)

	fmt.Println("Матрица после редукции:")
	fmt.Printf("%v\n", result)
}
