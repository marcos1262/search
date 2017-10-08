package main

import (
	"strconv"
	"math/rand"
	"time"
	"search/search"
)

type state struct {
	tab  [][]int
	zero [2]int
	cost int
}

func (s state) Cost() int {
	return s.cost
}

func (s state) IsGoal() bool {
	var n = len(s.tab)
	if s.tab[n-1][n-1] != 0 {
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s.tab[i][j] != i*n+j+1 && j+1 < n {
				return false
			}
		}
	}
	return true
}

func (s state) String() (str string) {
	str += "State [\n"
	for i, line := range s.tab {
		str += "\t"
		for j, item := range line {
			str += strconv.Itoa(item)
			if j < len(line)-1 {
				str += " "
			} else if i < len(line)-1 {
				str += "\n"
			}
		}
	}
	str += "\n\t" +
		"zero: (" +
		strconv.Itoa(s.zero[0]) + "," + strconv.Itoa(s.zero[1]) +
		")\n" +
		"\tcost: " + strconv.Itoa(s.cost) + "]"
	return
}

func (s state) Actions() []search.Action {
	var actions = []search.Action{}
	var n = len(s.tab)

	if s.zero[0] > 0 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]int{s.zero[0] - 1, s.zero[1]},
			})
	}
	if s.zero[1] > 0 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]int{s.zero[0], s.zero[1] - 1},
			})
	}
	if s.zero[0] < n-1 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]int{s.zero[0] + 1, s.zero[1]},
			})
	}
	if s.zero[1] < n-1 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]int{s.zero[0], s.zero[1] + 1},
			})
	}

	return actions
}

func (s state) Equals(other interface{}) bool {
	var state = other.(state)

	if s.cost != state.cost {
		return false
	}
	if s.zero != state.zero {
		return false
	}
	for i, line := range s.tab {
		for j := range line {
			if s.tab[i][j] != state.tab[i][j] {
				return false
			}
		}
	}
	return true
}

func (s state) Successor(act search.Action) search.State {
	var a = act.(action)
	var n = len(s.tab)
	var tab [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		tab[i] = make([]int, n)
		for j := 0; j < n; j++ {
			tab[i][j] = s.tab[i][j]
		}
	}

	var successor = state{zero: [2]int{}, cost: s.cost, tab: tab}

	successor.tab[a.zero[0]][a.zero[1]] = successor.tab[a.old[0]][a.old[1]]
	successor.tab[a.old[0]][a.old[1]] = 0
	successor.zero = a.old

	return successor
}

func NewGoalState(n int) search.State {
	var goal [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		goal[i] = make([]int, n)
		for j := 0; j < n; j++ {
			goal[i][j] = i*n + j + 1
		}
	}
	goal[n-1][n-1] = 0

	return state{
		tab:  goal,
		cost: 0,
		zero: [2]int{n - 1, n - 1},
	}
}

// Generate a random solvable initial state
func NewInitialState(n int) search.State {
	var initial [][]int = make([][]int, n)
	var random = rand.New(rand.NewSource(time.Now().UnixNano()))
	var unordered = random.Perm(n * n)
	var zero = [2]int{0, 0}

	for i := 0; i < n; i++ {
		initial[i] = make([]int, n)
		for j := 0; j < n; j++ {
			initial[i][j] = unordered[i*n+j]
			if initial[i][j] == 0 {
				zero[0] = i
				zero[1] = j
			}
		}
	}

	if isSolvable(initial) {
		println("OK")
		return state{
			tab:  initial,
			cost: 1,
			zero: zero,
		}
	} else {
		println("NOT")
		return NewInitialState(n)
	}

}
