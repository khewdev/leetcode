/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    result := &ListNode{0, head}
    left := result
    
    right := head
    
    for n > 0 && right != nil {
        right = right.Next
        n--
    }
    
    for right != nil {
        left = left.Next
        right = right.Next
    }
    
    left.Next = left.Next.Next
    
    return result.Next
}