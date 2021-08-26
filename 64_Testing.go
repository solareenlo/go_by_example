package Testing_test

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	res := IntMin(2, -2)
	if res != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", res)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			res := IntMin(tt.a, tt.b)
			if res != tt.want {
				t.Errorf("got %d, want %d", res, tt.want)
			}
		})
	}
}
