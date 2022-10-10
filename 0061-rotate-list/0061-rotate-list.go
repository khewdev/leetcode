/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    
    length, tail := 1, head
    for tail.Next != nil {
        tail = tail.Next
        length++
    }
    
    k = k % length
    if k == 0 {
        return head
    }
    
    cur := head
    for i:= 0; i < length - k - 1; i++ {
        cur = cur.Next
    }
    
    newHead := cur.Next
    cur.Next = nil
    tail.Next = head
    
    return newHead
}