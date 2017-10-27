package search

import (
	"strconv"
	"time"
	"fmt"
	"math"
	"search/data_structures"
)

var UseHeuristic bool

type Printable interface {
	String() (str string)
}

type State interface {
	Printable
	Cost() int
	Hash() string
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
	level         int
	meanBranching float32
}

func (s Statistic) String() (str string) {
	return "Statistic [" +
		"\n\t time: " + s.time.String() +
		"\n\t qtd. steps: " + strconv.Itoa(s.qtdSteps) +
		"\n\t level: " + strconv.Itoa(s.level) +
		"\n\t mean branching: " + fmt.Sprintf("%.2f", s.meanBranching) +
		"\n]"
}

func GraphSearch(problem Problem, border Border) (solution Solution, stats Statistic) {
	var start = time.Now()
	var explored = make(map[string]Node)

	border.Put(Node{problem.Initial(), nil, nil, 0, 1})

	for !border.Empty() {
		var node = border.Get().(Node)
		var actions = node.Actions()

		if node.IsGoal() {
			stats.time = time.Since(start)
			stats.level = node.level
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

func DFS_limited(problem Problem) (solution Solution, stats Statistic) {
	var start = time.Now()
	var explored = make(map[string]Node)
	var border = data_structures.NewStack()

	border.Put(Node{problem.Initial(), nil, nil, 0, 1})

	for !border.Empty() {
		var node = border.Get().(Node)
		var actions = node.Actions()

		if node.IsGoal() {
			stats.time = time.Since(start)
			stats.level = node.level
			solution = buildSolution(&node)
			return
		}

		explored[node.State.Hash()] = node
		stats.meanBranching = (stats.meanBranching*float32(stats.qtdSteps) +
			float32(len(actions))) / float32(stats.qtdSteps+1)
		stats.qtdSteps++

		if node.level >= 30 {
			continue
		}

		for _, action := range actions {
			var newNode = newNode(node, action)
			if isNewNode(newNode, explored, border) {
				border.Put(newNode)
			}
		}
	}

	return Solution{}, stats
}

func A_star(problem Problem) (solution Solution, stats Statistic) {
	var start = time.Now()
	var border = data_structures.NewPriorityQueue()

	border.Put(Node{problem.Initial(), nil, nil, 0, 1})

	for !border.Empty() {
		var node = border.Get().(Node)
		var actions = node.Actions()

		if node.IsGoal() {
			stats.time = time.Since(start)
			stats.level = node.level
			solution = buildSolution(&node)
			return
		}
		println("\n"+node.String()+"\n")

		stats.meanBranching = (stats.meanBranching*float32(stats.qtdSteps) +
			float32(len(actions))) / float32(stats.qtdSteps+1)
		stats.qtdSteps++

		for _, action := range actions {
			var newNode = newNode(node, action)
			println(newNode.String())
			border.Put(newNode)
		}
		time.Sleep(3*time.Second)
	}

	return Solution{}, stats
}

func IDA_star(problem Problem) (solution Solution, stats Statistic) {
	var start = time.Now()
	var limit = problem.Initial().Cost()
	var border = data_structures.NewPriorityQueue()

	border.Put(Node{problem.Initial(), nil, nil, 0, 1})

	for {
		for !border.Empty() {
			var node = border.Get().(Node)
			var actions= node.Actions()

			if node.Cost() > limit {
				limit = node.Cost()
				break
			}
			if node.IsGoal() {
				stats.level = node.level
				stats.time = time.Since(start)
				solution = buildSolution(&node)
				return
			}

			stats.meanBranching = (stats.meanBranching*float32(stats.qtdSteps) +
				float32(len(actions))) / float32(stats.qtdSteps+1)
			stats.qtdSteps++

			for _, action := range actions {
				var newNode= newNode(node, action)
				border.Put(newNode)
			}
		}

		if len(solution) > 0 {
			stats.time = time.Since(start)
			return solution, stats
		}
		if limit == math.MaxInt64 {
			return
		}
	}
}

//func search(node Node, stats Statistic, limit *int) (solution Solution) {
//	var nextLimit = math.MaxInt64
//	var actions = node.Actions()
//
//	if node.Cost() > *limit {
//		*limit = node.Cost()
//		return
//	}
//	if node.IsGoal() {
//		stats.level = node.level
//		return buildSolution(&node)
//	}
//
//	stats.meanBranching = (stats.meanBranching*float32(stats.qtdSteps) +
//		float32(len(actions))) / float32(stats.qtdSteps+1)
//	stats.qtdSteps++
//
//	for _, action := range actions {
//		var newNode = newNode(node, action)
//		solution = search(newNode, stats, limit)
//		if len(solution) > 0 {
//			return solution
//		} else{
//			nextLimit = int(math.Min(float64(*limit), float64(nextLimit)))
//			println(nextLimit)
//		}
//	}
//
//	*limit = nextLimit
//	return
//}

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
