package linkedlist

import (
	"testing"
)

func TestLinkedListAdd(t *testing.T) {
	l := NewLinkedList()

	tests := []struct {
		name string
		n    int
		want int
	}{
		{"", 1, 1},
		{"", 10, 2},
		{"", 20, 3},
		{"", 30, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Add(tt.n)
			if l.size != tt.want {
				t.Errorf("list.size = %v, want %v", l.size, tt.want)
			}
		})
	}
}

func TestLinkedListString(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	want := "1-> 2-> 3-> 4"
	got := l.String()
	if want != got {
		t.Errorf("list.size = %v, want %v", got, want)
	}
}
