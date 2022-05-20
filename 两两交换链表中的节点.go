package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	i := 0
	var pre *ListNode
	for head != nil {
		i++
		if i%2 == 0 {
			next := head.Next
			head.Next = pre
			pre, head = head, next
		} else {
			pre.Next = head.Next
		}
	}
	return pre
}
