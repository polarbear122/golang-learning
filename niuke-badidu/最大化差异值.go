package main

import (
	"fmt"
	"sort"
)

func main() {
	N := 0
	fmt.Scan(&N)
	a := []int{}
	for i := 0; i < N; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}
	b := []int{}
	for i := 0; i < N; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		b = append(b, tmp)
	}
	fmt.Printf("%v,%v", a, b)
	a1 := a
	b1 := b
	sort.Ints(a)
	sort.Ints(b)
	aMid := a[len(a)/2]
	bMid := b[len(b)/2]
	for i := 0; i < N; i++ {
		if a1[i] < aMid {
			if b1[i] < bMid {
				break
			}
		}
	}
}
