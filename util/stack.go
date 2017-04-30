package util

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
	if len(stack.Items) == 0 {
		return nil
	}

	item := stack.Items[0]
	stack.Items = append(stack.Items[:0], stack.Items[1:]...)
	stack.Count -= 1
	return item
}


