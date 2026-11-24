package main

import "testing"

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		m, n int
		want int64
	}{
		{3, 7, 28},
		{3, 2, 3},
		{1, 1, 1},
		{2, 2, 2},
		{1, 10, 1},
		{10, 10, 48620},
		{0, 5, 0},
	}
	for _, c := range cases {
		if got := UniquePaths(c.m, c.n); got != c.want {
			t.Errorf("UniquePaths(%d,%d) = %d, want %d", c.m, c.n, got, c.want)
		}
	}
	if UniquePaths(3, 7) != UniquePaths(7, 3) {
		t.Error("paths should be symmetric in m and n")
	}
}
