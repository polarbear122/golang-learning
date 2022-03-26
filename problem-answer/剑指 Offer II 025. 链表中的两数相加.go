package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type listNode struct {
 *     Val int
 *     Next *listNode
 * }
 */
func addTwoNumbers1(l1 *listNode, l2 *listNode) *listNode {
	l1 = ReverseChain(l1)
	l2 = ReverseChain(l2)
	cur1, cur2 := l1, l2
	l3 := &listNode{}
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
		cur.Next = &listNode{
			Val: remain,
		}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &listNode{
			Val: 1,
		}
	}
	return ReverseChain(l3.Next)
}

func ReverseChain(list *listNode) *listNode {
	if list == nil {
		return nil
	}
	var pre *listNode
	cur := list
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
