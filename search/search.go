package search

type State interface {
	Cost() int
	Actions() []Action
	Sucessor(action Action) State
	IsGoal() bool
	Equals(state State) bool
}

type Action interface {
}

type Problem interface {
	Initial() State
	Goals() []State
}

type Border interface {
	Put(node Node)
	Get() Node
	IsEmpty() bool
}

type Solution []Action

type Node struct {
	State
	previous State
	pathCost int
}

func newNode(previous Node, action Action) Node {
	var succ = previous.Sucessor(action)
	return Node{succ, previous, previous.pathCost + succ.Cost()}
}

func Search(problem Problem, border Border) (solution Solution, stepCount int) {
	var explored = []Node{}

	border.Put(Node{problem.Initial(), nil, 0})
	for {
		if border.IsEmpty() {
			return
		}
		var node = border.Get()
		stepCount++
		if node.IsGoal() {
			solution = buildSolution(node)
		}
		explored = append(explored, node)
		for action := range node.Actions() {
			var newNode = newNode(node, action)
			if isNewState(newNode, explored, border) {
				border.Put(newNode)
			}
		}
	}
}

// the state is not in the explored set or the border
func isNewState(state State, explored []Node, border Border) bool {
	for stateExplored := range explored {
		if state.Equals(stateExplored) {
			return false
		}
	}
}

func buildSolution(node Node) []Action {

}
