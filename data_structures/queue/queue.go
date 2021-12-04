package queue

type IntQueue struct {
	queue []int
}

func (q *IntQueue) Enqueue(elements ...int) {
	q.queue = append(q.queue, elements...)
}

func (q *IntQueue) Dequeue() int {
	lastQueuedItem := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]
	return lastQueuedItem
}
