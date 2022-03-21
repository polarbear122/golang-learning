package main

import "fmt"

type SortListNode struct {
	Val  int
	Next *SortListNode
}

func main() {
	node5 := SortListNode{
		Val: 23,
	}
	node4 := SortListNode{
		Val:  10,
		Next: &node5,
	}

	node3 := SortListNode{
		Val:  5,
		Next: &node4,
	}

	node2 := SortListNode{
		Val:  1,
		Next: &node3,
	}

	node1 := SortListNode{
		Val:  12,
		Next: &node2,
	}

	head := SortListNode{
		Val:  0,
		Next: &node1,
	}
	insertionSortList(&head)
	fmt.Println("test")

}

func insertionSortList(head *SortListNode) *SortListNode {
	if head == nil {
		return nil
	}
	dummyHead := &SortListNode{Next: head}
	lastSorted, nextNode := head, head.Next
	for nextNode != nil {
		if lastSorted.Val <= nextNode.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= nextNode.Val {
				prev = prev.Next
			}
			lastSorted.Next = nextNode.Next
			nextNode.Next = prev.Next
			prev.Next = nextNode
		}
		nextNode = lastSorted.Next
	}
	return dummyHead.Next
}
