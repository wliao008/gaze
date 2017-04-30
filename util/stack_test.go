package util

import (
	"testing"
	"time"
	"math/rand"
)

func TestPush(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(100)
	stack := Stack{}
	for i := 0; i < num; i++ {
		stack.Push(i)
	}

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

func BenchmarkPush(b *testing.B) {
	stack := Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}
