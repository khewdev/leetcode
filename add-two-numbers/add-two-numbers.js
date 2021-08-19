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
var addTwoNumbers = function(l1, l2) {
    let dummyHead = new ListNode(0);
    let curr = dummyHead;
    let carry = 0;
    
    while(l1 != null || l2 != null) {
        let x = (l1 != null) ? l1.val : 0;
        let y = (l2 != null) ? l2.val : 0;
        
        let sum = carry + x + y;
        carry = parseInt(sum / 10);
        lastDigit = sum % 10;
        curr.next = new ListNode(lastDigit);
        curr = curr.next;
        
        if(l1 != null) l1 = l1.next;
        if(l2 != null) l2 = l2.next;
    }
    
    if(carry > 0) {
        curr.next = new ListNode(carry);
    }
    
    return dummyHead.next;
    
};


// public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
//     ListNode dummyHead = new ListNode(0);
//     ListNode p = l1, q = l2, curr = dummyHead;
//     int carry = 0;
//     while (p != null || q != null) {
//         int x = (p != null) ? p.val : 0;
//         int y = (q != null) ? q.val : 0;
//         int sum = carry + x + y;
//         carry = sum / 10;
//         curr.next = new ListNode(sum % 10);
//         curr = curr.next;
//         if (p != null) p = p.next;
//         if (q != null) q = q.next;
//     }
//     if (carry > 0) {
//         curr.next = new ListNode(carry);
//     }
//     return dummyHead.next;
// }