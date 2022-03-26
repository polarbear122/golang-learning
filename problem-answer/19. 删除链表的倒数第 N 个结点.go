package main

import "fmt"

func main() {
	node5 := listNode{Val: 5}
	node4 := listNode{Val: 4, Next: &node5}
	node3 := listNode{Val: 3, Next: &node4}
	node2 := listNode{Val: 2, Next: &node3}
	node1 := listNode{Val: 1, Next: &node2}
	node0 := listNode{Val: 0, Next: &node1}
	fmt.Println(node0)
	list := removeNthFromEnd(&node0, 2)
	fmt.Println(list)
}

func removeNthFromEnd(list *listNode, N int) *listNode {
	listCopy := list
	length := 0
	for ; listCopy != nil; listCopy = listCopy.Next {
		length++
	}
	listHead := &listNode{
		Val:  0,
		Next: list,
	}
	listDelete := listHead
	for i := 0; i < length-N; i++ {
		listDelete = listDelete.Next
	}
	listDelete.Next = listDelete.Next.Next
	return listHead.Next
}
