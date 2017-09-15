package main

type State struct {
	title, target string
	cost          int
}

func (s State) Cost() int {
	return s.cost
}
func (s State) Actions() []Action {

}
func (s State) Sucessor(action Action) State {

}
func (s State) IsGoal() bool {

}
func (s State) Equals(state State) bool {
	return s.title == state.title && s.target == state.target
}

func main() {
}
