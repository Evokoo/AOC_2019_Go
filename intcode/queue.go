package intcode

// ========================
// QUEUE
// ========================
type Queue []int

func NewQueue() Queue {
	return []int{}
}
func (q *Queue) Push(value int) {
	*q = append(*q, value)
}
func (q *Queue) Pop() int {
	removed := (*q)[0]
	*q = (*q)[1:]
	return removed
}
