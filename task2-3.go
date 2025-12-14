package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func randMas(mas [][]int) {
	for i := 0; i < len(mas); i++ {
		for j := 0; j < len(mas[i]); j++ {
			mas[i][j] = rand.IntN(10)
		}
	}
}

func standMasMultiply(mas [][]int, mas2 [][]int) [][]int {
	n := len(mas[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for i := range len(mas) {
		for j := range len(mas2[0]) {
			for k := range len(mas[0]) {
				result[i][j] += mas[i][k] * mas2[k][j]
			}
		}
	}

	return result
}

func goMasMultiply(mas [][]int, mas2 [][]int) [][]int {
	n := len(mas[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	var wg sync.WaitGroup
	wg.Add(n * n)

	for i := range len(mas) {
		for j := range len(mas2[0]) {
			go func(i, j int) {
				defer wg.Done()

				for k := range len(mas[0]) {
					result[i][j] += mas[i][k] * mas2[k][j]
				}
			}(i, j)
		}
	}

	wg.Wait()

	return result
}

func matrixCompare(mas [][]int, mas2 [][]int) bool {
	n := len(mas[0])
	isEqual := true
	for i := range n {
		for j := range n {
			if mas[i][j] != mas2[i][j] {
				isEqual = false
			}
		}
	}

	return isEqual
}

func task2_3() {
	var n int
	fmt.Print("Введите число n для создания матрицы: ")
	fmt.Scan(&n)

	mas := make([][]int, n)
	for i := 0; i < n; i++ {
		mas[i] = make([]int, n)
	}
	randMas(mas)

	mas2 := make([][]int, n)
	for i := 0; i < n; i++ {
		mas2[i] = make([]int, n)
	}
	randMas(mas2)

	fmt.Println("Матрица 1:")
	fmt.Println(mas)
	fmt.Println("Матрица 2:")
	fmt.Println(mas2)

	startStand := time.Now()
	standartMulMatrix := standMasMultiply(mas, mas2)
	standTime := time.Since(startStand)

	startGo := time.Now()
	goMulMatrix := goMasMultiply(mas, mas2)
	goTime := time.Since(startGo)

	fmt.Println("Перемноженная матрица:")
	fmt.Println(standartMulMatrix)

	if matrixCompare(standartMulMatrix, goMulMatrix) {
		fmt.Println("OK")
		fmt.Printf("Последовательное умножение: %v\n", standTime)
		fmt.Printf("Параллельное умножение:    %v\n", goTime)
	} else {
		panic("Результаты умножений не совпадают")
	}

}
