package util

import (
	"testing"
	"time"
	"math/rand"
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
