package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var head *ListNode = createLinkedList(arr)
	printLinkedList(reverseListOptimal(head))

	arr = []int{1, 2}
	head = createLinkedList(arr)
	printLinkedList(reverseListOptimal(head))

	arr = []int{}
	head = createLinkedList(arr)
	printLinkedList(reverseListOptimal(head))

}

func reverseListBruteforce() {
	// Initially take a stack and pushing all nodes one after other into the stack
	// create a new linkedlist and add one after other poping out from the stack
	// Return the new linkedlist
}

func reverseListOptimal(head *ListNode) *ListNode {
	temp := head
	var prev *ListNode = nil
	for temp != nil {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return prev
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
