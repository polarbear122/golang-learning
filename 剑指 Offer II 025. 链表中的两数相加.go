package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = ReverseChain(l1)
	l2 = ReverseChain(l2)
	cur1, cur2 := l1, l2
	l3 := &ListNode{}
	cur := l3
	carry := 0
	for cur1 != nil || cur2 != nil {
		sum := 0
		if cur1 != nil {
			sum += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}
		remain := (sum + carry) % 10
		carry = (sum + carry) / 10
		cur.Next = &ListNode{
			Val: remain,
		}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &ListNode{
			Val: 1,
		}
	}
	return ReverseChain(l3.Next)
}

func ReverseChain(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}
	var pre *ListNode
	cur := list
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
