package stack

type Stack []int

func (s *Stack) Push(state int) {
	*s = append(*s, state)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	pop := (*s)[index]
	*s = (*s)[:index]

	return pop, true
}

func (s *Stack) PopMultiple(n int) ([]int, bool) {
	if len(*s) == 0 || len(*s) < n {
		return nil, false
	}

	var pop = make([]int, 0, 25)
	index := len(*s) - 1

	for i := 0; i < n; i++ {
		pop = append(pop, (*s)[index-i])
		*s = (*s)[:index-i]
	}

	return pop, true
}

func (s *Stack) PeekTop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	return (*s)[len(*s)-1], true
}

func (s *Stack) PeekDown(index int) (int, bool) {
	if len(*s) == 0 || len(*s) < index {
		return 0, false
	}

	baseIndex := len(*s) - 1

	return (*s)[baseIndex-index], true
}
