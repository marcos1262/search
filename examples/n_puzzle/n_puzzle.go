package main

import (
	"search/search"
	"strconv"
	"search/data_structures"
)

type problem struct {
	initial search.State
	goal    search.State
}

func (p problem) Goal() search.State {
	return p.goal
}
func (p problem) Initial() search.State {
	return p.initial
}
func (p problem) String() (str string) {
	return "Problem [" +
		"\n\tInitial State: "+p.initial.String()+
		"\n\tGoal State: "+p.goal.String()+
			"\n]"
}
func (p problem) Heuristic(state state) int {
	// TODO Manhattan distance
	//for i, line := range state.tab {
	//	for j := range line {
	//
	//	}
	//}
	return 0
}

type action struct {
	// move piece from old position to zero position
	old  [2]uint8
	zero [2]uint8
}

func (a action) String() (str string) {
	return "Move zero from " +
		"(" +
		strconv.Itoa(int(a.zero[0])) + "," + strconv.Itoa(int(a.zero[1])) +
		") to (" +
		strconv.Itoa(int(a.old[0])) + "," + strconv.Itoa(int(a.old[1])) +
		")"
}

func main() {
	var n = 3

	var problem = problem{
		goal:    NewGoalState(n),
		initial: NewInitialState(n),
	}

	println(problem.String())

	println("--------- Breath-First Search ---------")
	var border = data_structures.NewQueue()
	var solution, statistic = search.GraphSearch(problem, border)
	println(solution.String())
	println(statistic.String())

	println("--------- Depth-First Search ---------")
	var border2 = data_structures.NewStack()
	solution, statistic = search.GraphSearch(problem, border2)
	println(solution.String())
	println(statistic.String())

	//println("--------- Uniform Cost Search ---------")
	//var border3 = data_structures.NewPriorityQueue()
	//solution, statistic = search.GraphSearch(problem, border3)
	//println(solution.String())
	//println(statistic.String())


}
