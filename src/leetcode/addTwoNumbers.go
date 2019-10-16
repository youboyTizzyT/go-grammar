/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 18:02
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.

 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l = &ListNode{}
	pre := l
	flag := 0
	for l1 != nil || l2 != nil {
		pre.Next = &ListNode{}
		p := pre.Next
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		p.Val = (x + y + flag) % 10
		flag = (x + y + flag) / 10
		pre = p
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if flag != 0 {
		pre.Next = &ListNode{Val: flag}
	}
	return l.Next
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	l := result
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		l.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		l = l.Next
	}
	return result.Next
}
func main() {
	addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})
}
