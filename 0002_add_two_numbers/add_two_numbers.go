package main

import "fmt"

//
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

func printListNode(l *ListNode) {
	work := l
	for work != nil {
		fmt.Print(work.Val)
		if work.Next != nil {
			fmt.Print(" -> ")
		}
		work = work.Next
	}
	fmt.Println()
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	s := &ListNode{}
	w1, w2, w, c := l1, l2, s, 0

	for w1 != nil && w2 != nil {
		w.Next = &ListNode{}
		w = w.Next
		w.Val = w1.Val + w2.Val + c
		c = w.Val / 10
		w.Val = w.Val % 10

		w1 = w1.Next
		w2 = w2.Next
	}

	if w1 != nil {
		for w1 != nil {
			w.Next = &ListNode{}
			w = w.Next
			w.Val = w1.Val + c
			c = 0

			w1 = w1.Next
		}
	} else if w2 != nil {
		for w2 != nil {
			w.Next = &ListNode{}
			w = w.Next
			w.Val = w2.Val + c
			c = 0
		}
	}

	if c != 0 {
		w.Next = &ListNode{c, nil}
	}

	return s.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, q, w := l1, l2, res
	var x, y, carry int
	for p != nil || q != nil {
		if p == nil {
			x = 0
		} else {
			x = p.Val
		}
		if q == nil {
			y = 0
		} else {
			y = q.Val
		}
		summary := carry + x + y
		carry = summary / 10
		w.Next = &ListNode{summary % 10, nil}
		w = w.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		w.Next = &ListNode{carry, nil}
	}
	return res.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, q, w, sum := l1, l2, res, 0
	for p != nil || q != nil {
		sum /= 10
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		w.Next = &ListNode{sum % 10, nil}
		w = w.Next
	}
	if sum/10 == 1 {
		w.Next = &ListNode{1, nil}
	}
	return res.Next
}

func main() {
	l1 := ListNode{5, nil}

	l2 := ListNode{5, nil}

	printListNode(addTwoNumbers(&l1, &l2))
}
