package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	n1 := 2
	head1 := createList(arr1)
	head1 = removeNthFromEnd(head1, n1)
	printList(head1)

	arr2 := []int{1}
	n2 := 1
	head2 := createList(arr2)
	head2 = removeNthFromEnd(head2, n2)
	printList(head2)

	arr3 := []int{1, 2}
	n3 := 2
	head3 := createList(arr3)
	head3 = removeNthFromEnd(head3, n3)
	printList(head3)

	arr4 := []int{1, 2, 3}
	n4 := 1
	head4 := createList(arr4)
	head4 = removeNthFromEnd(head4, n4)
	printList(head4)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}


func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, v := range arr[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

