package gaze

import (
	"math/rand"
	"testing"
	"time"
)

func TestPush(t *testing.T) {
	stack, num := createStack()
	if stack.Count != num {
		t.Errorf("Push() %d items, got %d", num, stack.Count)
	}
}

func TestPop(t *testing.T) {
	stack := Stack{}
	item := stack.Pop()
	if item != nil {
		t.Errorf("Pop() from empty stack, got %v", item)
	}
}

func TestPop2(t *testing.T) {
	stack := Stack{}
	item := stack.Pop()
	if item != nil {
		t.Errorf("Pop() from empty stack, got %v", item)
	}
}

func TestPop3(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	item := stack.Pop()
	if item != 2 {
		t.Errorf("Pop(), should get 2, got %v", item)
	}

	item2 := stack.Pop()
	if item2 != 1 {
		t.Errorf("Pop(), should get 1, got %v", item2)
	}

	if stack.Count != 0 {
		t.Errorf("Pop(), count should be 0, go %d", stack.Count)
	}
}

func TestIsEmpty(t *testing.T) {
	stack, _ := createStack()
	for stack.Count > 0 {
		stack.Pop()
	}

	result := stack.IsEmpty()
	if !result {
		t.Errorf("IsEmpty() should be true, got %v", result)
	}
}

func TestPeek(t *testing.T) {
	stack, num := createStack()
	stack.Peek()
	if stack.Count != num {
		t.Errorf("Peek() should not remove item. Had %d, now %d", num, stack.Count)
	}
}

func TestPeek2(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	item := stack.Peek()
	if item != 2 {
		t.Errorf("Pop(), should get 2, got %v", item)
	}
}

func BenchmarkPush(b *testing.B) {
	stack := Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func createStack() (Stack, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(100) + 1
	stack := Stack{}
	for i := 0; i < num; i++ {
		stack.Push(i)
	}
	return stack, num
}
