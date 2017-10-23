package main

import (
	"strconv"
	"math/rand"
	"time"
	"search/search"
	"math"
)

type state struct {
	tab  [][]uint8
	zero [2]uint8
	hash uint64
	heuristic uint8
}

func (s state) Cost() int {
	var cost = 1
	if search.UseHeuristic {
		cost += int(s.heuristic)
	}
	return cost
}

func (s state) Hash() uint64 {
	return s.hash
}

func (s state) IsGoal() bool {
	var n = len(s.tab)
	if s.tab[n-1][n-1] != 0 {
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if int(s.tab[i][j]) != i*n+j+1 && j+1 < n {
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
			str += strconv.Itoa(int(item))
			if j < len(line)-1 {
				str += " "
			} else if i < len(line)-1 {
				str += "\n"
			}
		}
	}
	str += "\n\t" +
		"zero: (" +
		strconv.Itoa(int(s.zero[0])) + "," + strconv.Itoa(int(s.zero[1])) +
		")]"
	return
}

func (s state) Actions() []search.Action {
	var actions = []search.Action{}
	var n = len(s.tab)

	if s.zero[0] > 0 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]uint8{s.zero[0] - 1, s.zero[1]},
			})
	}
	if s.zero[1] > 0 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]uint8{s.zero[0], s.zero[1] - 1},
			})
	}
	if int(s.zero[0]) < n-1 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]uint8{s.zero[0] + 1, s.zero[1]},
			})
	}
	if int(s.zero[1]) < n-1 {
		actions = append(actions,
			action{zero: s.zero,
				old: [2]uint8{s.zero[0], s.zero[1] + 1},
			})
	}

	return actions
}

func (s state) Equals(other interface{}) bool {
	var state = other.(state)
	return s.hash == state.hash
}

func (s state) Successor(act search.Action) search.State {
	var a = act.(action)
	var n = len(s.tab)
	var tab [][]uint8 = make([][]uint8, n)

	for i := 0; i < n; i++ {
		tab[i] = make([]uint8, n)
		for j := 0; j < n; j++ {
			tab[i][j] = s.tab[i][j]
		}
	}

	var successor = state{
		zero: [2]uint8{},
		tab:  tab}

	successor.tab[a.zero[0]][a.zero[1]] = successor.tab[a.old[0]][a.old[1]]
	successor.tab[a.old[0]][a.old[1]] = 0

	successor.zero = a.old
	successor.hash = hash(tab)
	successor.heuristic = manhattan(tab)

	return successor
}

func hash(tab [][]uint8) uint64 {
	var hash uint64
	var n = len(tab)

	for i, line := range tab {
		for j, item := range line {
			hash += uint64(item) * uint64(math.Pow10(i*n+j))
		}
	}
	return hash
}

func manhattan(tab [][]uint8) uint8 {
	// Manhattan distance
	var n = len(tab)
	var dist uint8

	for i, line := range tab {
		for j, item := range line {
			var iGoal, jGoal = int(item-1)/n, int(item-1)%n
			if item == 0 {
				iGoal, jGoal = 2, 2
			}
			var di = uint8(math.Abs(float64(i - iGoal)))
			var dj = uint8(math.Abs(float64(j - jGoal)))
			dist += di + dj
		}
	}
	return dist
}

func NewGoalState(n int) search.State {
	var goal = make([][]uint8, n)

	for i := 0; i < n; i++ {
		goal[i] = make([]uint8, n)
		for j := 0; j < n; j++ {
			goal[i][j] = uint8(i*n + j + 1)
		}
	}
	goal[n-1][n-1] = 0

	return state{
		tab:  goal,
		zero: [2]uint8{uint8(n - 1), uint8(n - 1)},
		hash: hash(goal),
		heuristic:manhattan(goal),
	}
}

// Generate a random solvable initial state
func NewInitialState(n int) search.State {
	var initial [][]uint8 = make([][]uint8, n)
	var zero = [2]uint8{0, 0}

	var random = rand.New(rand.NewSource(time.Now().UnixNano()))
	var unordered = random.Perm(n * n)

	for i := 0; i < n; i++ {
		initial[i] = make([]uint8, n)
		for j := 0; j < n; j++ {
			initial[i][j] = uint8(unordered[i*n+j])
			if initial[i][j] == 0 {
				zero[0] = uint8(i)
				zero[1] = uint8(j)
			}
		}
	}

	if isSolvable(initial) {
		println("Solvable")
		return state{
			tab:  initial,
			zero: zero,
			hash: hash(initial),
			heuristic:manhattan(initial),
		}
	} else {
		println("Not Solvable")
		return NewInitialState(n)
	}

}
