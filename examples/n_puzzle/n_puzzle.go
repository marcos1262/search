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
		"\nInitial State: "+p.initial.String()+
		"\nGoal State: "+p.goal.String()+
			"\n]"
}
func (p problem) Heuristic(state state) int {
	// TODO Manhattan distance
	return 0
}

type action struct {
	// move piece from old position to zero position
	old  [2]int
	zero [2]int
}

func (a action) String() (str string) {
	return "Move zero from " +
		"(" +
		strconv.Itoa(a.zero[0]) + "," + strconv.Itoa(a.zero[1]) +
		") to (" +
		strconv.Itoa(a.old[0]) + "," + strconv.Itoa(a.old[1]) +
		")"
}

func main() {
	var n = 3

	var problem = problem{
		goal:    NewGoalState(n),
		initial: NewInitialState(n),
	}

	var border = data_structures.NewQueue()

	var solution, steps = search.BlindSearch(problem, border)

	println(problem.String())
	println(solution.String(), "\nstep count:", steps)
}
