package day06

import (
	"fmt"
	"strconv"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	switch part {
	case aoc.Part1:
		result := part1(lines)
		return strconv.Itoa(result), nil
	case aoc.Part2:
		result := part2(lines)
		return strconv.Itoa(result), nil
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

func part1(lines []string) int {
	var total int
	group := map[rune]bool{}
	for _, line := range lines {
		for _, c := range line {
			group[c] = true
		}
		if line == "" {
			total += len(group)
			group = map[rune]bool{}
		}
	}
	total += len(group)
	return total
}

func part2(lines []string) int {
	var total int

	groupSize := 0
	answerCounts := map[rune]int{}
	for _, line := range lines {
		if line == "" {
			for _, v := range answerCounts {
				if v == groupSize {
					total++
				}
			}
			answerCounts = map[rune]int{}
			groupSize = 0
			continue
		}
		groupSize++
		for _, c := range line {
			answerCounts[c]++
		}
	}
	for _, v := range answerCounts {
		if v == groupSize {
			total++
		}
	}
	return total
}
