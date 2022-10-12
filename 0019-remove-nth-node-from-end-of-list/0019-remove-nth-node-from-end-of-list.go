/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    
    left := dummy
    right := head
    
    for n > 0 {
        right = right.Next
        n--
    }
    
    for right != nil {
        left = left.Next
        right = right.Next
    }
    
    left.Next = left.Next.Next
    
    return dummy.Next
}