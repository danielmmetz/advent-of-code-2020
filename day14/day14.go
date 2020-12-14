package day14

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

var memAssignmentExpr = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
var setBitmaskExpr = regexp.MustCompile(`mask = ([01X]{36})`)

func MainE(lines []string, part aoc.Part) (string, error) {
	switch part {
	case aoc.Part1, aoc.Part2:
		instructions, err := parse(lines)
		if err != nil {
			return "", fmt.Errorf("error parsing program: %w", err)
		}
		result := evaluate(instructions, part)
		return strconv.Itoa(result), nil
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

func parse(lines []string) ([]instruction, error) {
	var instructions []instruction
	for i, line := range lines {
		if match := setBitmaskExpr.FindStringSubmatch(line); match != nil {
			instructions = append(instructions, instruction{kind: setBitmask, bitmask: bitmask(match[1])})
			continue
		}
		if match := memAssignmentExpr.FindStringSubmatch(line); match != nil {
			location, _ := strconv.Atoi(match[1])
			value, _ := strconv.Atoi(match[2])
			instructions = append(instructions, instruction{kind: memoryAssignment, assignment: assignment{location: location, value: value}})
			continue
		}
		return instructions, fmt.Errorf("error parsing line %d: no instruction pattern match: %s", i, line)
	}
	return instructions, nil
}

func evaluate(instructions []instruction, part aoc.Part) int {
	var memory map[int]int
	switch part {
	case aoc.Part1:
		memory = evaluate1(instructions)
	case aoc.Part2:
		memory = evaluate2(instructions)
	}
	var total int
	for _, value := range memory {
		total += value
	}
	return total
}

func evaluate1(instructions []instruction) map[int]int {
	memory := make(map[int]int)
	var b bitmask

	for _, instruction := range instructions {
		switch instruction.kind {
		case setBitmask:
			b = instruction.bitmask
		case memoryAssignment:
			memory[instruction.assignment.location] = b.render1(instruction.assignment.value)
		}
	}
	return memory
}

func evaluate2(instructions []instruction) map[int]int {
	memory := make(map[int]int)
	var b bitmask

	for _, instruction := range instructions {
		switch instruction.kind {
		case setBitmask:
			b = instruction.bitmask
		case memoryAssignment:
			locations := b.render2(instruction.assignment.location)
			for _, location := range locations {
				memory[location] = instruction.assignment.value
			}
		}
	}
	return memory
}

type instruction struct {
	kind       instructionKind
	bitmask    bitmask
	assignment assignment
}

type instructionKind string

const (
	setBitmask       instructionKind = "setBitmask"
	memoryAssignment instructionKind = "memoryAssignment"
)

type bitmask string

func (b bitmask) render1(value int) int {
	var result int
	for i := 0; i < 36; i++ {
		switch b[35-i] {
		case zero:
			continue
		case one:
			result += 1 << i
		case x:
			result += ((value >> i) & 1) << i
		}
	}
	return result
}

func (b bitmask) render2(value int) []int {
	results := []int{value}
	for bitPos := 0; bitPos < 36; bitPos++ {
		switch b[35-bitPos] {
		case zero:
			continue
		case one:
			for i := range results {
				if bit := results[i] >> bitPos & 1; bit != 1 {
					results[i] += 1 << bitPos
				}
			}
		case x:
			var newValues []int
			for i := range results {
				bit := results[i] >> bitPos & 1
				switch bit {
				case 0:
					newValues = append(newValues, results[i]+1<<bitPos)
				case 1:
					newValues = append(newValues, results[i]-1<<bitPos)
				}
			}
			results = append(results, newValues...)
		}
	}
	sort.Ints(results)
	return results
}

const (
	zero = '0'
	one  = '1'
	x    = 'X'
)

type assignment struct {
	location int
	value    int
}
