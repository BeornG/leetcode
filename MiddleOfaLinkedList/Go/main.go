package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// Example 1
	// Input: [1,2,3,4,5]
	// Output: [3,4,5]

	example1 := &ListNode{Val: 1}
	example1.Next = &ListNode{Val: 2}
	example1.Next.Next = &ListNode{Val: 3}
	example1.Next.Next.Next = &ListNode{Val: 4}
	example1.Next.Next.Next.Next = &ListNode{Val: 5}

	middle := middleNode(example1)

	for middle != nil {
		fmt.Print(middle.Val, " ")
		middle = middle.Next
	}
	fmt.Println()

}

func middleNode(head *ListNode) *ListNode {

	// Floyd’s Cycle Finding Algorithm
	// https://www.geeksforgeeks.org/floyds-cycle-finding-algorithm/
	slow := head // slow pointer moves one step at a time
	fast := head // fast pointer moves two steps at a time

	// when fast reaches the end of the list, slow will be at the middle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}
