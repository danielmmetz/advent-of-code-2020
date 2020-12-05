package day05

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	var seatIDs []int
	for _, line := range lines {
		id := boardingPass(line).SeatID()
		seatIDs = append(seatIDs, id)
	}

	switch part {
	case aoc.Part1:
		var maxSeatID int
		for _, id := range seatIDs {
			if id > maxSeatID {
				maxSeatID = id
			}
		}
		return strconv.Itoa(maxSeatID), nil
	case aoc.Part2:
		sort.Ints(seatIDs)
		for i := 0; i < len(seatIDs)-1; i++ {
			if seatIDs[i+1]-seatIDs[i] == 2 {
				return strconv.Itoa(seatIDs[i] + 1), nil
			}
		}
		return "", fmt.Errorf("unable to determine seat id")
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

type boardingPass string

func (p boardingPass) Row() int {
	lower, upper := 0, 127
	for i := 0; i < 7; i++ {
		delta := 1 << (6 - i)
		switch p[i] {
		case 'F':
			upper -= delta
		case 'B':
			lower += delta
		}
	}
	return lower
}

func (p boardingPass) Column() int {
	lower, upper := 0, 7
	for i := 0; i < 3; i++ {
		delta := 1 << (2 - i)
		switch p[7+i] {
		case 'L':
			upper -= delta
		case 'R':
			lower += delta
		}
	}
	return lower
}

func (p boardingPass) SeatID() int {
	return p.Row()*8 + p.Column()
}
