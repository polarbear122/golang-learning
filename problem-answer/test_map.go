package main

import (
	"fmt"
	"time"
)

type Job struct {
	id int
}

var m map[int]map[string]*Job

func main() {
	for {
		var j = Job{id: 1}
		mm := map[string]*Job{"ok": &j}
		//m=map[10]mm
		fmt.Println(mm)
		m = map[int]map[string]*Job{10: mm}
		fmt.Println(m)
		<-time.NewTimer(1 * time.Second).C
	}
}
