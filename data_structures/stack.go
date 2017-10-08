package data_structures

type Stack struct {
	items []interface{}
}

func (s Stack) Len() int {
	return len(s.items)
}

func (s Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Get() interface{} {
	var item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Put(item interface{}) {
	s.items = append(s.items, item)
}

func (s Stack) Contains(comparable interface {
	Equals(other interface{}) bool
}) bool {
	for _, n := range s.items {
		if comparable.Equals(n) {
			return true
		}
	}
	return false
}

func NewStack() *Stack {
	return &Stack{[]interface{}{}}
}
