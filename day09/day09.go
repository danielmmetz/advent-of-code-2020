package day09

import (
	"fmt"
	"strconv"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part, opts ...option) (string, error) {
	preambleLength := 25
	for _, opt := range opts {
		preambleLength = opt()
	}
	var numbers []int
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("error converting to number: %w", err)
		}
		numbers = append(numbers, number)
	}
	switch part {
	case aoc.Part1:
		result, err := firstViolation(numbers, preambleLength)
		return strconv.Itoa(result), err
	case aoc.Part2:
		target, err := firstViolation(numbers, preambleLength)
		if err != nil {
			return "", fmt.Errorf("no violation found to target for contiguous number search: %w", err)
		}
		smallest, largest, err := contiguousComponents(numbers, target)
		if err != nil {
			return "", fmt.Errorf("no contigous stretch found summing to target value: %w", err)
		}
		return strconv.Itoa(smallest + largest), nil
	default:
		return "", fmt.Errorf("invalid part")
	}
}

type option func() int

func preambleLength(i int) option {
	return func() int { return i }
}

func firstViolation(numbers []int, preambleLength int) (int, error) {
	if len(numbers) < preambleLength {
		return 0, fmt.Errorf("no volations found")
	}
	rm := newRingMap(preambleLength)
	for i := 0; i < preambleLength; i++ {
		rm.Add(numbers[i])
	}
	for _, candidate := range numbers[preambleLength:] {
		var propertySatisfied bool
		for i := 0; i < preambleLength-1 && !propertySatisfied; i++ {
			for j := i + 1; j < preambleLength; j++ {
				if rm.Entries()[i] == rm.Entries()[j] {
					continue
				}
				if rm.Entries()[i]+rm.Entries()[j] == candidate {
					propertySatisfied = true
					break
				}
			}
		}
		if !propertySatisfied {
			return candidate, nil
		}
		rm.Add(candidate)
	}
	return 0, fmt.Errorf("no volations found")
}

func contiguousComponents(numbers []int, target int) (int, int, error) {
	if len(numbers) < 2 {
		return 0, 0, fmt.Errorf("input list too small for contiguous component search")
	}
	left, right, sum := 0, 1, numbers[0]+numbers[1]
	for right < len(numbers) {
		switch {
		case sum == target:
			min, max := numbers[left], numbers[right]
			for i := left; i <= right; i++ {
				if numbers[i] < min {
					min = numbers[i]
				}
				if numbers[i] > max {
					max = numbers[i]
				}
			}
			return min, max, nil
		case sum < target, left == right:
			if right+1 >= len(numbers) {
				return 0, 0, fmt.Errorf("exhausted inputs without finding target contiguous stretch")
			}
			right++
			sum += numbers[right]
		case sum > target:
			sum -= numbers[left]
			left++
		}
	}
	return 0, 0, fmt.Errorf("exhausted inputs without finding target contiguous stretch")
}

type ringMap struct {
	contains map[int]int
	ring     []int
	insertAt int
}

func newRingMap(size int) ringMap {
	return ringMap{
		contains: make(map[int]int),
		ring:     make([]int, size),
	}
}

func (r *ringMap) Add(i int) {
	old := r.ring[r.insertAt]
	r.contains[old]--
	if r.contains[old] == 0 {
		delete(r.contains, old)
	}

	r.contains[i]++
	r.ring[r.insertAt] = i
	r.insertAt = (r.insertAt + 1) % len(r.ring)
}

func (r *ringMap) Entries() []int {
	return r.ring
}
