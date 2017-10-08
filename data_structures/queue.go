package data_structures

import (
	"github.com/Workiva/go-datastructures/queue"
)

type Queue struct {
	*queue.Queue
}

func (q Queue) Len() int {
	return int(q.Queue.Len())
}

func (q *Queue) Get() interface{} {
	var items, _ = q.Queue.Get(1)
	return items[0]
}

func (q *Queue) Put(item interface{}) {
	q.Queue.Put(item)
}

func NewQueue() *Queue {
	return &Queue{queue.New(100)}
}
