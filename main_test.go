package main

import "testing"

func TestPass(t *testing.T) {
	what := sum(4, 5)
	if what != 9 {
		t.Fail()
	}

}

func TestFail(t *testing.T) {
	what := sum(2, 5)
	if what != 9 {
		t.Errorf("Sum does not equal 9")
	}

}

func TestSumm(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 5},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
