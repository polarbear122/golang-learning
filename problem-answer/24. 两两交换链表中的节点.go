package main

import "fmt"

func main() {
	n := Create([]int{1, 2, 3, 4, 5})
	Print(swapPairs(n))
}

func Print(root *listNode) {
	for root != nil {
		fmt.Printf("%v ", root.Val)
		root = root.Next
	}
}

type listNode struct {
	Val  int
	Next *listNode
}

func swapPairs(head *listNode) *listNode {
	dummyHead := &listNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

func Create(array []int) *listNode {
	root := new(listNode)
	cur := root
	for i := 0; i < len(array); i++ {
		tmp := new(listNode)
		tmp.Val = array[i]
		cur.Next = tmp
		cur = cur.Next
	}
	return root.Next
}
