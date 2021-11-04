package nav

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nojnilsson/aoc/puzzleio"
)

type Coordinate struct {
	X int
	Y int
}

func intAbs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func (c1 *Coordinate) Add(c2 Coordinate) {
	c1.X = c1.X + c2.X
	c1.Y = c1.X + c2.Y
	fmt.Println(*c1)
}

func (c1 Coordinate) ManhattanDistance(c2 Coordinate) int {
	xDelta := intAbs(c2.X - c1.X)
	yDelta := intAbs(c2.Y - c1.Y)
	return xDelta + yDelta
}

func FindExtremes(coordinates []Coordinate) (Coordinate, Coordinate) {
	var maxX int = coordinates[0].X
	var minX int = coordinates[0].X
	var maxY int = coordinates[0].Y
	var minY int = coordinates[0].Y

	for _, c := range coordinates[1:] {
		if c.X > maxX {
			maxX = c.X
		}
		if c.X < minX {
			minX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		if c.Y < minY {
			minY = c.Y
		}
	}
	return Coordinate{minX, minY}, Coordinate{maxX, maxY}
}

// Should move to aocinput?
func ReadCoordinates(fname string) []Coordinate {
	var coordinates []Coordinate
	lines := puzzleio.ReadLines(fname)
	for _, l := range lines {
		ais := strings.Split(l, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(ais[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(ais[1]))
		c := Coordinate{x, y}
		coordinates = append(coordinates, c)
	}
	return coordinates
}
