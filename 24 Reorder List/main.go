package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var head *ListNode = createLinkedList(arr)
	reorderList(head)
	printLinkedList(head)

	arr = []int{1, 2}
	head = createLinkedList(arr)
	reorderList(head)
	printLinkedList(head)

	arr = []int{1, 2, 3, 4}
	head = createLinkedList(arr)
	reorderList(head)
	printLinkedList(head)

}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := slow.Next
	slow.Next = nil

	secondHalf = reverse(secondHalf)

	merge(head, secondHalf)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func merge(first *ListNode, second *ListNode) {
	for first != nil && second != nil {
		temp1 := first.Next
		temp2 := second.Next

		first.Next = second
		second.Next = temp1

		first = temp1
		second = temp2
	}
}

// Function to create a linked list from a slice
func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Function to print the linked list
func printLinkedList(head *ListNode) {
	current := head
	fmt.Print("[ ")
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Print("]\n")
}
