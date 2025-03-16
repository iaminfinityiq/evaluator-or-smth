package helpers

type IntSet struct {
	Set map[int]bool
}

func (s IntSet) Add(value int) {
	s.Set[value] = true
}

func (s IntSet) Remove(value int) {
	s.Set[value] = false
}

func (s IntSet) Contains(value int) bool {
	return s.Set[value]
}
