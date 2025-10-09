package main

import (
	"fmt"
	"sort"
)

const n2 = 4

func Generate(items [n2]string) func() ([n2]string, bool) {
	a := items
	nextPerm := func(p *[n2]string) bool {
		i := len(p) - 2
		for i >= 0 && !(p[i] < p[i+1]) {
			i--
		}
		if i < 0 {
			return false
		}

		j := len(p) - 1
		for !(p[i] < p[j]) {
			j--
		}

		p[i], p[j] = p[j], p[i]

		for l, r := i+1, len(p)-1; l < r; l, r = l+1, r-1 {
			p[l], p[r] = p[r], p[l]
		}
		return true
	}

	return func() ([n2]string, bool) {
		if nextPerm(&a) {
			return a, true
		}
		var empty [n2]string
		return empty, false
	}
}

func task9() {
	p := [n2]string{"a", "b", "c", "d"}
	s := p[:]
	sort.Strings(s)
	tst := Generate(p)
	for ok := true; ok; p, ok = tst() {
		fmt.Println(p)
	}
}
