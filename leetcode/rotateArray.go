package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateArray() {
	arr := []int{1, 2, 3, 4, 5}
	rotate := 2
	res := rotateArray(arr, rotate)
	fmt.Println(res)

	listNode, k := &ListNode{
		Val:  1,
		Next: nil,
	}, 2
	res2 := rotateRight(listNode, k)
	fmt.Println(res2)
}

func rotateArray(arr []int, k int) []int {
	idx := k%len(arr) + 1

	tmp := arr[idx:]
	arr = append(tmp, arr[:idx]...)

	return arr
}

func rotateRight(head *ListNode, k int) *ListNode {
	// If the list is empty or has only one node, return it as is
	if head == nil || head.Next == nil {
		return head
	}

	// If k is zero, the list doesn't need to be rotated
	if k == 0 {
		return head
	}

	// Find the last node
	n, last := 1, head
	for last.Next != nil {
		last = last.Next
		n++
	}

	// Compute the number of rotations needed
	k %= n

	// Find the new tail node (the node at position n - k)
	newTail := head
	for i := 0; i < n-k-1; i++ {
		newTail = newTail.Next
	}

	// The new head is the node following the new tail
	newHead := newTail.Next

	// Set the next pointer of the new tail to nil
	newTail.Next = nil

	// Set the next pointer of the last node to the original head
	last.Next = head

	return newHead
}
