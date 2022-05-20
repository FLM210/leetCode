package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head.Next != nil {
		next := head.Next
		head.Next = pre
		pre, head = head, next
	}
	return pre
}
