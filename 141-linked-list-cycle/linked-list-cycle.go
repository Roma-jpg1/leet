/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]int)

    if head == nil {
        return false
    }

    for head.Next != nil{
        m[head]++
        if m[head] == 2{
            return true
        }
        head = head.Next
    }
    return false
    
}