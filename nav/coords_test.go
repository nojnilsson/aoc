package nav

import "testing"

func TestIntAbs(t *testing.T) {
	testTuples := []struct {
		n int
		a int
	}{
		{0, 0},
		{-1, 1},
		{3, 3},
		{-3, 3},
	}
	for _, tuple := range testTuples {
		abs := intAbs(tuple.n)
		if abs != tuple.a {
			t.Errorf("intAbs(%d) error. Expected: %d, got: %d", tuple.n, tuple.a, abs)
		}
	}
}

func TestManhattanDistance(t *testing.T) {
	testTuples := []struct {
		c1 Coordinate
		c2 Coordinate
		d  int
	}{
		{
			Coordinate{X: 0, Y: 0},
			Coordinate{X: 1, Y: 1},
			2,
		},
		{
			Coordinate{X: 0, Y: 0},
			Coordinate{X: 3, Y: 3},
			6,
		},
		{
			Coordinate{X: 0, Y: 0},
			Coordinate{X: -2, Y: -2},
			4,
		},
	}
	for _, tuple := range testTuples {
		distance := tuple.c1.ManhattanDistance(tuple.c2)
		if distance != tuple.d {
			t.Error("Error calculating ManhattanDistance from", tuple.c1,
				" to", tuple.c2, "Expected:", tuple.d, " got:", distance)
		}
	}
}

func TestFindExtremes(t *testing.T) {
	testTuples := []struct {
		cs       []Coordinate
		origin   Coordinate
		opposite Coordinate
	}{
		{
			[]Coordinate{{0, 0}, {1, 1}, {2, 2}},
			Coordinate{0, 0},
			Coordinate{2, 2},
		},
		{
			[]Coordinate{{0, 0}, {1, 1}, {2, 2}, {-3, 0}, {0, 7}},
			Coordinate{-3, 0},
			Coordinate{2, 7},
		},
	}

	for _, tuple := range testTuples {
		ori, opp := FindExtremes(tuple.cs)
		if ori != tuple.origin || opp != tuple.opposite {
			t.Error("Error finding extremes: expected ", tuple.origin, tuple.opposite,
				" Got:", ori, opp)
		}
	}
}
