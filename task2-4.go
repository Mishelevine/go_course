package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(fname string) <-chan string {
	ch := make(chan string, 5)
	go func() {
		defer close(ch)
		file, err := os.Open(fname)
		if err != nil {
			log.Println("Ошибка открытия файла " + fname)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()
	return ch
}

func getNames(lines <-chan string) <-chan string {
	out := make(chan string, 5)

	go func() {
		defer close(out)
		for line := range lines {
			if strings.HasPrefix(line, "From ") {
				fields := strings.Fields(line)
				if len(fields) > 1 {
					out <- fields[1]
				}
			}
		}
	}()

	return out
}

func collectNames(names <-chan string) map[string]int {
	counts := make(map[string]int)
	for name := range names {
		counts[name]++
	}
	return counts
}

func task2_4() {
	senders := collectNames(getNames(readFile("mbox.txt")))

	fmt.Println("Отправители:")
	for email, count := range senders {
		fmt.Printf("%s: %d\n", email, count)
	}

	var maxSender string
	maxCount := 0
	for sender, cnt := range senders {
		if cnt > maxCount {
			maxCount = cnt
			maxSender = sender
		}
	}

	fmt.Printf("Больше всего писем у %s - %d\n", maxSender, maxCount)
}
