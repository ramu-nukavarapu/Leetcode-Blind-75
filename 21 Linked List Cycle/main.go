package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{3, 2, 0, -4}
	pos := 1
	var head *ListNode = createLinkedList(arr, pos)
	fmt.Println(hasCycleOptimal(head))

	arr = []int{1, 2}
	pos = 0
	head = createLinkedList(arr, pos)
	fmt.Println(hasCycleOptimal(head))

	arr = []int{1}
	pos = -1
	head = createLinkedList(arr, pos)
	fmt.Println(hasCycleOptimal(head))

}

func hasCycleBruteforce() {
	// Initially take a array/stack and pushing all nodes address one after other into the stack
	// while pushing check once while the address is already in the stack if yes return true
	// else return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycleOptimal(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// Function to create a linked list from a slice
func createLinkedList(arr []int, pos int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	var targetNode *ListNode // This will store the node at the `pos` position

	// Create the linked list and keep track of the target node if pos >= 0
	for i, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next

		// If the current index matches the position, store the node
		if i == pos-1 {
			targetNode = current
		}
	}

	// If pos == 0, the last node should point to the head
	if pos == 0 {
		targetNode = head
	}

	// Make the last node point to the target node if pos is valid
	if pos >= 0 {
		current.Next = targetNode
	}

	return head
}
