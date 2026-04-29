/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var pr *ListNode
    cur := head

    for cur != nil{
        next:=cur.Next
        cur.Next=pr
        pr = cur
        cur = next
    }
    return pr
    
}