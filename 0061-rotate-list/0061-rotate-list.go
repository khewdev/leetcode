/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    
    length, tail := 1, head
    
    for tail.Next != nil {
        length++
        tail = tail.Next
    }
    
    k = k % length
    if k == 0 {
        return head
    }
    
    currentNode := head
    for i := 0; i < length - k - 1; i++ {
        currentNode = currentNode.Next
    }
    
    newHead := currentNode.Next
    currentNode.Next = nil
    tail.Next = head
    
    return newHead
}