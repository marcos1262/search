package search

import (
	"strconv"
	"github.com/Workiva/go-datastructures/queue"
)

type Node struct {
	State
	previous *Node
	action   Action
	pathCost int
	level    int
}

func (n Node) String() (str string) {
	str += "Node [\n" + n.State.String()
	if n.action != nil {
		str += "\naction: " + n.action.String()
	}
	str += "\ncost: " + strconv.Itoa(n.Cost()) +
		"\npathCost: " + strconv.Itoa(n.pathCost) +
		"\nlevel: " + strconv.Itoa(n.level) +
		"]"
	return
}

func (n Node) Equals(other interface{}) bool {
	var node = other.(Node)
	return n.State.Equals(node.State)
}

func (n Node) Compare(other queue.Item) int {
	if n.Cost() > other.(Node).Cost() {
		return 1
	} else if n.Cost() < other.(Node).Cost() {
		return -1
	} else {
		return 0
	}
}

func newNode(previous Node, action Action) Node {
	var successor = previous.Successor(action)
	return Node{
		successor,
		&previous,
		action,
		previous.pathCost + successor.Cost(),
		previous.level + 1,
	}
}

// Verify if a node is not in the explored set or the border
func isNewNode(node Node, explored map[string]Node, border Border) bool {
	_, ok := explored[node.Hash()]
	return !ok && !border.Contains(node)
}
