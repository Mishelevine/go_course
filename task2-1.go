package main

import "fmt"

func task2_1() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)

	ch := make(chan struct{}, 2)
	defer close(ch)

	go func() {

		for i := 0; i <= n; i++ {
			if i%2 == 0 {
				fmt.Printf("%v\n", i)
			}
		}

		ch <- struct{}{}
	}()

	go func() {

		for i := 0; i <= n; i++ {
			if i%2 == 1 {
				fmt.Printf("%v\n", i)
			}
		}

		ch <- struct{}{}
	}()

	<-ch
	<-ch
}
