package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	rs, err := parse(lines)
	if err != nil {
		return "", fmt.Errorf("error parsing map: %w", err)
	}
	switch part {
	case aoc.Part1:
		return strconv.Itoa(rs.TreesEncountered(3, 1)), nil
	case aoc.Part2:
		result := 1
		slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
		for _, slope := range slopes {
			result *= rs.TreesEncountered(slope[0], slope[1])
		}
		return strconv.Itoa(result), nil
	default:
		return "", fmt.Errorf("error: invalid part: %v", part)
	}
}

func parse(lines []string) (rows, error) {
	var rs rows
	for _, line := range lines {
		var r []bool
		for _, c := range line {
			switch c {
			case '.':
				r = append(r, false)
			case '#':
				r = append(r, true)
			default:
				return rs, fmt.Errorf("unknown entity encountered in map: %v", c)
			}
		}
		rs = append(rs, row{isTree: r})
	}
	return rs, nil
}

type rows []row

func (r rows) String() string {
	result := make([]string, len(r))
	for i, row := range r {
		result[i] = row.String()
	}
	return strings.Join(result, "\n")
}

func (r rows) TreesEncountered(rightStepSize, downStepSize int) int {
	var collisions int
	for x, y := 0, 0; y < len(r); x, y = x+rightStepSize, y+downStepSize {
		if r[y].IsTree(x) {
			collisions++
		}
	}
	return collisions
}

type row struct {
	isTree []bool
}

func (r row) String() string {
	result := make([]string, len(r.isTree))
	for i, c := range r.isTree {
		if c {
			result[i] = "#"
		} else {
			result[i] = "."
		}
	}
	return strings.Join(result, "")
}

func (r row) IsTree(col int) bool {
	return r.isTree[col%len(r.isTree)]
}
