type MyQueue struct {
    in  []int
    out []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (q *MyQueue) pushToOut() {
    for len(q.in) > 0 {
        n := len(q.in) - 1
        q.out = append(q.out, q.in[n])
        q.in = q.in[:n]
    }
}

func (q *MyQueue) Push(x int) {
    q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
    if len(q.out) == 0 {
        q.pushToOut()
    }
    n := len(q.out) - 1
    val := q.out[n]
    q.out = q.out[:n]
    return val
}

func (q *MyQueue) Peek() int {
    if len(q.out) == 0 {
        q.pushToOut()
    }
    return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
    return len(q.in) == 0 && len(q.out) == 0
}