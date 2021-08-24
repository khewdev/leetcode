/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
    let temp_node = new ListNode(0);
    let current_node = temp_node;
    
    while(l1 !== null && l2 !== null) {
        if(l1.val < l2.val) {
            current_node.next = l1;
            l1 = l1.next;
        } else {
            current_node.next = l2;
            l2 = l2.next;
        }
        current_node = current_node.next;
    }
    
    if(l1 !== null) {
        current_node.next = l1;
        current_node = current_node.next;
    }
    
    if(l2 !== null) {
        current_node.next = l2;
        current_node = current_node.next;
    }
    
    return temp_node.next;
};
