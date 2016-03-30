package util

import (
	"testing"
)

func TestRound(t *testing.T) {
	cases := []struct {
		in  float64
		exp int
	}{
		{0, 0},
		{0.49, 0},
		{0.5, 1},
		{1, 1},
		{4.49, 4},
		{4.5, 5},
		{5, 5},
	}

	for i, c := range cases {
		if out := Round(c.in); out != c.exp {
			t.Errorf("Case %d: expected %d, but got %d", i, c.exp, out)
		}
	}
}
