package stack

// Slices are simply just too powerful in go not to use
// for any collective datastructure bar linked structures
// As such we are using a slice.
type Stack struct {
	store []interface{}
}

func (s *Stack) Pop() interface{} {
	popped := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return popped
}

// Append an item(s) to the stack
func (s *Stack) Push(data ...interface{}) {
	s.store = append(s.store, data...)
}

// Count it
func (s *Stack) Count() int {
	return len(s.store)
}
