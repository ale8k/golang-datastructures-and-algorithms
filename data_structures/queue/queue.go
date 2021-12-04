package queue

import "errors"

type IntQueue struct {
	queue []int
}

func (q *IntQueue) Enqueue(elements ...int) {
	q.queue = append(q.queue, elements...)
}

func (q *IntQueue) Dequeue() (int, error) {
	if len(q.queue) >= 1 {
		lastQueuedItem := q.queue[len(q.queue)-1]
		q.queue = q.queue[:len(q.queue)-1]
		return lastQueuedItem, nil
	}
	return 0, errors.New("cannot dequeue, queue has no queued items")
}
