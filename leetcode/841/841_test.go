package main

import (
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		in  [][]int
		out bool
	}{
		{in: [][]int{{1}, {2}, {3}, {}, {4}}, out: false},
		{in: [][]int{{1}, {2}, {3}, {}}, out: true},
		{in: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, out: false},
		{in: [][]int{{2, 3}, {}, {2}, {1, 3, 1}}, out: true},
	}

	for i, test := range tests {
		got := canVisitAllRooms(test.in)
		if got != test.out {
			t.Errorf("test %d: expected %t, got %t", i, test.out, got)
		}
	}
}
