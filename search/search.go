package search

import (
	"strconv"
	"time"
	"fmt"
)

type Printable interface {
	String() (str string)
}

type State interface {
	Printable
	Cost() int
	Hash() uint64
	IsGoal() bool
	Actions() []Action
	Equals(other interface{}) bool
	Successor(action Action) State
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

type Statistic struct {
	time          time.Duration
	qtdSteps      int
	meanBranching float32
}

func (s Statistic) String() (str string) {
	return "Statistic [" +
		"\n\t time: " + s.time.String() +
		"\n\t qtd. steps: " + strconv.Itoa(s.qtdSteps) +
		"\n\t mean branching:" + fmt.Sprintf("%.2f", s.meanBranching) +
		"\n]"
}

func GraphSearch(problem Problem, border Border) (solution Solution, stats Statistic) {
	var start = time.Now()
	var explored = make(map[uint64]Node)

	border.Put(Node{problem.Initial(), nil, nil, 0, 1})

	for !border.Empty() {
		var node = border.Get().(Node)
		var actions = node.Actions()

		if node.IsGoal() {
			stats.time = time.Since(start)
			solution = buildSolution(&node)
			return
		}

		explored[node.State.Hash()] = node
		stats.meanBranching = (stats.meanBranching*float32(stats.qtdSteps) +
			float32(len(actions))) / float32(stats.qtdSteps+1)
		stats.qtdSteps++

		for _, action := range actions {
			var newNode = newNode(node, action)
			if isNewNode(newNode, explored, border) {
				border.Put(newNode)
			}
		}
	}

	return Solution{}, stats
}

// Get actions selected by the search method
func buildSolution(node *Node) []Action {
	var actions = []Action{}
	for node.level > 1 {
		actions = append(actions, node.action)
		//println(node.State.String())
		node = node.previous
	}
	return actions
}
