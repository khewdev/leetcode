/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    
    for head != nil {
        nextNode := head.Next
        head.Next = prev
        prev = head
        head = nextNode
    }
    
    return prev
}