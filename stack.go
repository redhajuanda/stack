package stack

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() {
	if len(*s) != 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s Stack) Top() interface{} {
	return s[len(s)-1]
}

func (s Stack) Empty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}
