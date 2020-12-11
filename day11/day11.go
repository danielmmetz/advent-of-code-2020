package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	var g grid
	for r, line := range lines {
		g.rows = append(g.rows, make([]element, len(line)))
		for c, spot := range line {
			switch element(spot) {
			case floor, empty, occupied:
				g.rows[r][c] = element(spot)
			default:
				return "", fmt.Errorf("invalid element in grid at spot (%d, %d): %v", r, c, spot)
			}
		}
	}
	switch part {
	case aoc.Part1, aoc.Part2:
		var count int
		f, err := fixedPoint(g, part)
		for r := range f.rows {
			for _, v := range f.rows[r] {
				if v == occupied {
					count++
				}
			}
		}
		return strconv.Itoa(count), err
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

type element rune

const (
	floor    element = '.'
	empty    element = 'L'
	occupied element = '#'
)

const (
	stepLimit = 1000000
)

func fixedPoint(g grid, part aoc.Part) (grid, error) {
	for attempt := 0; attempt < stepLimit; attempt++ {
		next := g.next(part)
		if !equal(g, next) {
			g = next
			continue
		}
		return next, nil
	}
	return g, fmt.Errorf("no fixed point found after %d iterations", stepLimit)
}

func equal(g1, g2 grid) bool {
	if len(g1.rows) != len(g2.rows) {
		return false
	}
	for r := range g1.rows {
		if len(g1.rows[r]) != len(g2.rows[r]) {
			return false
		}
		for c := range g1.rows[r] {
			if g1.rows[r][c] != g2.rows[r][c] {
				return false
			}
		}
	}
	return true
}

type grid struct {
	rows [][]element
}

func (g grid) String() string {
	var b strings.Builder
	for r := range g.rows {
		for _, element := range g.rows[r] {
			b.WriteRune(rune(element))
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (g grid) adjacenciesOccupied1(r, c int) int {
	var hits int
	for y := r - 1; y <= r+1; y++ {
		for x := c - 1; x <= c+1; x++ {
			switch {
			case y == r && x == c:
				continue
			case y < 0 || y >= len(g.rows):
				continue
			case x < 0 || x >= len(g.rows[y]):
				continue
			case g.rows[y][x] == occupied:
				hits++
			}
		}
	}
	return hits
}
func (g grid) adjacenciesOccupied2(r, c int) int {
	vectors := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	var hits int
	for _, v := range vectors {
		y, x := r, c
		for done := false; !done; {
			y, x = y+v[0], x+v[1]
			switch {
			case y < 0 || y >= len(g.rows):
				done = true
			case x < 0 || x >= len(g.rows[y]):
				done = true
			case g.rows[y][x] == occupied:
				hits++
				done = true
			case g.rows[y][x] == empty:
				done = true
			}
		}
	}
	return hits
}

func (g grid) next(part aoc.Part) grid {
	if part == aoc.Part1 {
		return g.next1()
	}
	return g.next2()
}

func (g grid) clone() grid {
	clone := grid{rows: make([][]element, len(g.rows))}
	for i := range g.rows {
		clone.rows[i] = make([]element, len(g.rows[i]))
		copy(clone.rows[i], g.rows[i])
	}
	return clone
}

func (g grid) next1() grid {
	clone := g.clone()
	for r, row := range clone.rows {
		for c := range row {
			spot := g.rows[r][c]
			switch {
			case spot == empty && g.adjacenciesOccupied1(r, c) == 0:
				clone.rows[r][c] = occupied
			case spot == occupied && g.adjacenciesOccupied1(r, c) >= 4:
				clone.rows[r][c] = empty
			}
		}
	}
	return clone
}

func (g grid) next2() grid {
	clone := g.clone()
	for r, row := range clone.rows {
		for c := range row {
			spot := g.rows[r][c]
			switch {
			case spot == empty && g.adjacenciesOccupied2(r, c) == 0:
				clone.rows[r][c] = occupied
			case spot == occupied && g.adjacenciesOccupied2(r, c) >= 5:
				clone.rows[r][c] = empty
			}
		}
	}
	return clone
}
