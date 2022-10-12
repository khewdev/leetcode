/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    result := &ListNode{0, head}
    prev := result
    curr := head
    
    for curr != nil && curr.Next != nil {
        nextPair := curr.Next.Next
        second := curr.Next
        
        second.Next = curr
        curr.Next = nextPair
        prev.Next = second
        
        prev = curr
        curr = nextPair
    }
    
    return result.Next
}