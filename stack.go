package main

type Stack []interface{}

func (s Stack) Push(item interface{}) {
	s = append(s, item)
}
func (s Stack) Pop() interface{} {
	item := s[len(s)-1]
	s = s[:len(s)-1]
	return item
}
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
func (s Stack) Size() int {
	return len(s)
}
func (s Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func main() {
	problemStack := new(Stack)
	i := 0
	problemStack.Push(i)
	for len(problemStack) < 10 {
		problem, _ := problemStack.Top()
		pass := selectAnswer(resultState, problemStack.Size()-1, problem)
		if pass {
			problemStack.Push(0)
		} else {
			if problem.(int) < 3 {
				problemStack.Pop()
				answer := problem.(int) + 1
				problemStack.Push(answer)
			} else {
				problemStack.Pop()
				frontAnswer := problemStack.Pop()
				answer := frontAnswer + 1
				problemStack.Push(answer)
			}
		}
	}
}
