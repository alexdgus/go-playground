package main

func main()
{
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

func reverse(head *ListNode, count int) (newHead *ListNode) {
    newHead = head
    while newHead.next != nil {
        newHead = newHead.next
    }
    tracker := newHead
    while count > 1 {
        traversal := head
        for i := 0; i < count - 1; i++ {
            traversal = traversal.next
        }
        temp := traversal.next
        traversal.next = nil
    }
    while newHead.next != nil {
        newHead = newHead.next
    }
    
}