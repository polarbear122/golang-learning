package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello undo redo world undo redo undo redo undo undo good"
	fmt.Println(strDo(str))
}

const (
	undo = "undo"
	redo = "redo"
)

func strDo(str string) (modifyStr string) {
	arr := strings.Fields(str)
	tmpSta := stack{}
	for _, v := range arr {
		tmpSta.Put(v)
	}
	needRemove := 0
	for tmpSta.count >= 1 {
		tmpChar := tmpSta.list[tmpSta.count-1]
		tmpSta.count--
		if tmpChar != undo && tmpChar != redo {
			if needRemove == 0 {
				modifyStr = tmpChar + " " + modifyStr
			} else if needRemove > 0 {
				needRemove--
			}
		} else if tmpChar == redo {
			needRemove--
		} else {
			needRemove++
		}
	}
	return modifyStr
}

type stack struct {
	list  []string
	count int
}

func (s *stack) Pop() string {
	popNum := s.list[s.count-1]
	s.list = s.list[:s.count-1]
	return popNum
}

func (s *stack) Put(num string) {
	s.list = append(s.list, num)
	s.count++
}
