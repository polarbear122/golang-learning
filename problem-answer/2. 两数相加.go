package main

import "fmt"

// TwoListNode Definition for singly-linked list.
type TwoListNode struct {
	Val  int
	Next *TwoListNode
}

func addTwoNumbers(l1 *TwoListNode, l2 *TwoListNode) (result *TwoListNode) {
	var temp = result
	carry := 0
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if result == nil {
			fmt.Println(result)
			result = &TwoListNode{Val: sum}
			temp = result
		} else {
			temp.Next = &TwoListNode{Val: sum}
			temp = temp.Next
		}
	}
	if carry == 1 {
		temp.Next = &TwoListNode{Val: carry}
	}
	return result
}
func main() {

}
