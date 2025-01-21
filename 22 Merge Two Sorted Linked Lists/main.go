package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	var list1 *ListNode = createLinkedList(arr1)
	arr2 := []int{1, 2}
	var list2 *ListNode = createLinkedList(arr2)
	printLinkedList(mergeTwoLists(list1, list2))

	arr1 = []int{1, 7}
	list1 = createLinkedList(arr1)
	arr2 = []int{1, 2, 3, 14, 25}
	list2 = createLinkedList(arr2)
	printLinkedList(mergeTwoLists(list1, list2))

	arr1 = []int{}
	list1 = createLinkedList(arr1)
	arr2 = []int{1, 2, 3, 4, 5}
	list2 = createLinkedList(arr2)
	printLinkedList(mergeTwoLists(list1, list2))

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode = new(ListNode)
	temp := result
	for list1 != nil && list2 != nil {
		var node *ListNode = new(ListNode)
		if list1.Val > list2.Val {
			node.Val = list2.Val
			list2 = list2.Next
		} else {
			node.Val = list1.Val
			list1 = list1.Next
		}
		temp.Next = node
		temp = node
	}

	if list1 != nil {
		temp.Next = list1
	}

	if list2 != nil {
		temp.Next = list2
	}
	return result.Next
}

// Function to create a linked list from a slice
func createLinkedList(arr1 []int) *ListNode {
	if len(arr1) == 0 {
		return nil
	}

	list1 := &ListNode{Val: arr1[0]}
	current := list1
	for _, val := range arr1[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return list1
}

// Function to print the linked list
func printLinkedList(list1 *ListNode) {
	current := list1
	fmt.Print("[ ")
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Print("]\n")
}
