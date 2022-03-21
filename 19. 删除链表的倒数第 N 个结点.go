package main

import "fmt"

func main() {
	node5 := ListNode{Val: 5}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	node0 := ListNode{Val: 0, Next: &node1}
	fmt.Println(node0)
	list := removeNthFromEnd(&node0, 2)
	fmt.Println(list)
}

func removeNthFromEnd(list *ListNode, N int) *ListNode {
	listCopy := list
	length := 0
	for ; listCopy != nil; listCopy = listCopy.Next {
		length++
	}
	listHead := &ListNode{
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
