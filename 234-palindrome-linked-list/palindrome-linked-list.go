/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    second := reverseList(middleNode(head))
    first := head

    for second != nil {
        if first.Val != second.Val {
            return false
        }

        first = first.Next
        second = second.Next
    }

    return true    
}

func middleNode(head *ListNode) *ListNode {
    fast := head
    slow := head

    for fast != nil && fast.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
    }

    return slow

    
}

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