package main

import "fmt"

func main() {
	c3 := ListNode{Val: 5, Next: nil}
	c2 := ListNode{Val: 4, Next: &c3}
	c1 := ListNode{Val: 8, Next: &c2}
	b3 := ListNode{Val: 1, Next: &c1}
	b2 := ListNode{Val: 6, Next: &b3}
	a2 := ListNode{Val: 1, Next: &c1}
	headB := ListNode{Val: 5, Next: &b2}
	headA := ListNode{Val: 4, Next: &a2}
	fmt.Println(getIntersectionNode(&headA, &headB))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointers := make(map[*ListNode]int)
	a := headA
	b := headB

	for a != nil {
		pointers[a] = a.Val
		a = a.Next
	}

	for b != nil {
		if _, ok := pointers[b]; ok {
			return b
		}
		b = b.Next
	}

	return nil
}
