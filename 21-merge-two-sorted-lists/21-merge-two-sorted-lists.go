/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var resultNode = new(ListNode)
    var temp = resultNode
    
    for l1 != nil && l2 != nil {
        if(l1.Val < l2.Val) {
            temp.Next = l1
            l1 = l1.Next
        } else {
            temp.Next = l2
            l2 = l2.Next
        }
        temp = temp.Next
    }
    
    if l1 != nil {
        temp.Next = l1
    } else {
        temp.Next = l2
    }
    
    return resultNode.Next
}