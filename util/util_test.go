package util

import "testing"

func TestShuffle(t *testing.T) {
	var data []interface{}
	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)
	Shuffle(data)
	if data[0] == 1 || data[1] == 2 || data[2] == 3 {
		t.Errorf("Shuffle([1,2,3]), want items shuffled, got %v", data)
	}
}

func BenchmarkShuffle(b *testing.B) {
	var data []interface{}
	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)
	for i := 0; i < b.N; i++ {
		Shuffle(data)
	}
}
