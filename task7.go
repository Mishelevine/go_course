package main

import (
	"fmt"
	"sort"
)

func sortLines(a [][]float64) {
	for i := 0; i < len(a); i++ {
		sort.Slice(a[i], func(j, k int) bool { return a[i][j] > a[i][k] })
	}

	fmt.Printf("Отсортированый массив:\n%v\n", a)
}

func task7() {
	slices := [][]float64{
		{3, 1, 2},
		{4, 14, 10, 1},
	}

	fmt.Printf("Изначальная матрица:\n%v\n", slices)

	sortLines(slices)
}
