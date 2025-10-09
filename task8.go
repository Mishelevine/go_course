package main

import (
	"fmt"
	"sort"
)

func intersection(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	i, j := 0, 0
	res := make([]int, 0)

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func task8() {
	a := []int{1, 2, 2, 3, 4}
	b := []int{2, 2, 4, 5}

	fmt.Printf("Изначальный слайс a:\n%v\n", a)
	fmt.Printf("Изначальный слайс b:\n%v\n", b)

	fmt.Print("Слайс пересечение:\n")
	fmt.Println(intersection(a, b))
}
