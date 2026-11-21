package Utils

type Stack []interface{}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Pop() interface{} {
	h := *s
	l := len(h)
	el := h[l-1]
	*s = h[:l-1]
	return el
}

func (s *Stack) Length() int {
	return len(*s)
}

func NewStack() *Stack {
	return &Stack{}
}
