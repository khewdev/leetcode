/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var resultNode = new(ListNode)
    var tempNode = resultNode
    
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tempNode.Next = l1
            l1 = l1.Next
        } else {
            tempNode.Next = l2
            l2 = l2.Next
        }
        tempNode = tempNode.Next
    }
    
    if l1 != nil {
        tempNode.Next = l1
    } else {
        tempNode.Next = l2
    }
    
    return resultNode.Next
}