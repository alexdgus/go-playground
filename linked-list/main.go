package main

import "fmt"

// Linked list node type
type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	head := new(ListNode)
	head.val = 5
	head.next = nil
	tail := head
	for i := 0; i < 5; i++ {
		tail.next = new(ListNode)
		tail = tail.next
		tail.val = i
		tail.next = nil
	}
	curr := head
	for curr != nil {
		fmt.Println(curr)
		curr = curr.next
	}

	revHead := reverse(head)
	curr = revHead
	for curr != nil {
		fmt.Println(curr)
		curr = curr.next
	}
}

 func reorderList(head *ListNode)  {
    // Empty list, nothing to reorder
    if head == nil {
        return
    }

    traversalPointer := head
    count := 1
    while head.Next != nil {
        count++
        head = head.Next
    }

    // Only one element in the list, nothing to reorder
    if count < 2 {
        return
    }

    backCount := count / 2
    frontCount := count / 2
    if count % 2 == 1 {
        frountCount++
    }

    frontPtr := head
    backPtr := head
    for i := 0; i < frontCount - 1; i++ {
        backPtr = backPtr.next
    }
    tempPtr := backPtr.next
    backPtr.next = nil
    backPtr := tempPtr

    frontSwapPtr = head
    while count > 0 {
        traversalPointer = head
        for i := 0; i < count - 1; i++ {
            traversalPointer = traversalPointer->next
        }
        tempFrontNodePtr = frontSwapPtr.next
        frontSwapPtr.next = .next
    }
}

func reverse(head *ListNode) (revHead *ListNode) {
	// Set the last node as the head of the reversed list
	revHead = head
	for revHead.next != nil {
		revHead = revHead.next
	}
	revHead.next = nil

	revCurr := revHead
	var prev *ListNode
	var curr *ListNode

	for revCurr != nil {
		prev = nil
		curr = head
		// Traverse to the last node of list
		for curr.next != nil {
			prev = curr
			curr = curr.next
		}
		// Append it to the reversed list
		revCurr.next = curr
		revCurr = revCurr.next
		if prev != nil {
			// Remove the last node from the list
			prev.next = nil
		} else {
			// All nodes moved to new list
			break
		}
	}
	return revHead
}
