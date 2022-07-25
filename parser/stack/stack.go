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
