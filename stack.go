package gaze

type Stack struct {
	Count int
	Items []interface{}
}

var stack []interface{}

func (stack *Stack) Push(item interface{}) {
	stack.Items = append(stack.Items, item)
	stack.Count += 1
}

func (stack *Stack) Pop() interface{} {
	if stack.Count == 0 {
		return nil
	}

	item := stack.Items[stack.Count-1]
	stack.Items = stack.Items[:stack.Count-1]
	stack.Count -= 1
	return item
}

func (stack *Stack) Peek() interface{} {
	if stack.Count == 0 {
		return nil
	}

	return stack.Items[stack.Count-1]
}

func (stack *Stack) IsEmpty() bool {
	return stack.Count == 0
}
