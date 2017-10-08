package data_structures

import (
	"github.com/Workiva/go-datastructures/queue"
)

type PriorityQueue struct {
	*queue.PriorityQueue
}

func (p PriorityQueue) Size() int {
	return p.Len()
}

func (p PriorityQueue) IsEmpty() bool {
	return p.Empty()
}

func (p *PriorityQueue) Get() interface{} {
	var items, _ = p.PriorityQueue.Get(1)
	return items[0]
}

func (p *PriorityQueue) Put(item interface{}) {
	p.PriorityQueue.Put(item.(queue.Item))
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{queue.NewPriorityQueue(100, true)}
}
