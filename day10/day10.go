package day10

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	var joltages []int
	for _, line := range lines {
		joltage, err := strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("unable to parse joltage: %w", err)
		}
		joltages = append(joltages, joltage)
	}
	sort.Ints(joltages)
	switch part {
	case aoc.Part1:
		result, err := differences(joltages)
		return strconv.Itoa(result.one * result.three), err
	case aoc.Part2:
		result := combinations(joltages)
		return strconv.Itoa(result), nil
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

type diffCount struct {
	one, two, three int
}

func differences(joltages []int) (diffCount, error) {
	var current int
	var counts diffCount
	for _, j := range joltages {
		switch j - current {
		case 0:
			continue
		case 1:
			counts.one++
		case 2:
			counts.two++
		case 3:
			counts.three++
		default:
			return counts, fmt.Errorf("joltage jump too big: attempted to go from %d to %d", current, j)
		}
		current = j
	}
	counts.three++ // account for device adapter
	return counts, nil
}

func combinations(joltages []int) int {
	if len(joltages) == 0 {
		return 1
	}
	return combinationsR(append([]int{0}, joltages...), joltages[len(joltages)-1]+3, map[int]int{})
}

func combinationsR(joltages []int, current int, cache map[int]int) int {
	if count, ok := cache[current]; ok {
		return count
	}
	if len(joltages) == 0 {
		return 1
	}

	var successes int
	for skip := 1; skip <= 3; skip++ {
		if len(joltages) >= skip && current-joltages[len(joltages)-skip] <= 3 {
			successes += combinationsR(joltages[:len(joltages)-skip], joltages[len(joltages)-skip], cache)
		}
	}
	cache[current] = successes
	return successes
}
