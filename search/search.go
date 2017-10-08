package search

import (
	"strconv"
	"github.com/Workiva/go-datastructures/queue"
)

type State interface {
	Printable
	Comparable
	Cost() int
	IsGoal() bool
	Actions() []Action
	Successor(action Action) State
}

type Printable interface {
	String() (str string)
}

type Comparable interface {
	Equals(other interface{}) bool
}

type Action interface {
	Printable
}

type Problem interface {
	Printable
	Initial() State
}

type Border interface {
	Get() interface{}
	Len() int
	Put(interface{})
	Empty() bool
	Contains(comparable interface {
		Equals(other interface{}) bool
	}) bool
}

type Solution []Action

func (s Solution) String() (str string) {
	str += "Solution [\n"
	for i := len(s); i > 0; i-- {
		str += strconv.Itoa(len(s)-i+1) + ": " + s[i-1].String() + "\n"
	}
	return str + "]"
}

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
	str += "\npathCost: " + strconv.Itoa(n.pathCost) +
		"\nlevel: " + strconv.Itoa(n.level) +
		"]"
	return
}
func (n Node) Equals(other interface{}) bool {
	var node = other.(Node)
	return n.State.Equals(node.State)
}
func (n Node) Compare(other queue.Item) int {
	if n.pathCost > other.(Node).pathCost {
		return 1
	} else if n.pathCost < other.(Node).pathCost {
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

func BlindSearch(problem Problem, border Border) (solution Solution, stepCount int) {
	var explored = []Node{}

	border.Put(Node{problem.Initial(), nil, nil, 0, 1})

	for {
		if border.Empty() {
			break
		}
		var node = border.Get().(Node)
		//println(node.String())
		stepCount++
		if node.IsGoal() {
			solution = buildSolution(&node)
			return
		}
		explored = append(explored, node)
		for _, action := range node.Actions() {
			var newNode = newNode(node, action)
			if isNewState(newNode, explored, border) {
				//println(newNode.String())
				border.Put(newNode)
			}
		}
	}
	return Solution{}, stepCount
}

// Verify if a state is not in the explored set or the border
func isNewState(node Node, explored []Node, border Border) bool {
	if border.Contains(node) {
		return false
	}
	for _, nodeExplored := range explored {
		if node.Equals(nodeExplored) {
			return false
		}
	}
	return true
}

// Get actions selected by the search method
func buildSolution(node *Node) []Action {
	var actions = []Action{}
	for node.level > 1 {
		actions = append(actions, node.action)
		println(node.State.String())
		node = node.previous
	}
	return actions
}
