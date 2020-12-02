package aoc

import "fmt"

type Part int

const (
	Part1 Part = iota
	Part2
)

func ToPart(i int) (Part, error) {
	switch i {
	case 1:
		return Part1, nil
	case 2:
		return Part2, nil
	default:
		return Part(0), fmt.Errorf("invalid part specified: %d", i)
	}
}
