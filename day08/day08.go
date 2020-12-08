package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	var program []instruction
	for _, line := range lines {
		ins, err := parse(line)
		if err != nil {
			return "", fmt.Errorf("error parsing instruction: %w", err)
		}
		program = append(program, ins)
	}
	switch part {
	case aoc.Part1:
		result, err := evaluateUntilRepeat(program)
		if err != nil {
			return "", fmt.Errorf("error evaluating program: %w", err)
		}
		return strconv.Itoa(result), nil
	case aoc.Part2:
		for i := 0; i < len(program); i++ {
			mp, err := modified(program, i)
			if err != nil {
				continue
			}
			result, err := execute(mp)
			if err != nil {
				continue
			}
			return strconv.Itoa(result), nil
		}
		return "", fmt.Errorf("unable to identify corrupted line: input exhuasted")
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

func modified(program []instruction, i int) ([]instruction, error) {
	if i < 0 || i >= len(program) {
		return nil, fmt.Errorf("mutation specified for invalid instruction number")
	}
	clone := make([]instruction, len(program))
	copy(clone, program)
	switch clone[i].op {
	case nop:
		clone[i].op = jmp
	case jmp:
		clone[i].op = nop
	default:
		return clone, fmt.Errorf("modification requested for invalid operation: %s", clone[i].op)
	}
	return clone, nil
}

func evaluateUntilRepeat(program []instruction) (int, error) {
	value := 0
	seen := map[int]bool{}
	for i := 0; !seen[i]; {
		if i < 0 || i >= len(program) {
			return value, fmt.Errorf("error: tried to evaluate invalid instruction number: %d", i)
		}
		seen[i] = true
		switch program[i].op {
		case nop:
			i++
		case acc:
			value += program[i].arg
			i++
		case jmp:
			i += program[i].arg
		default:
			return value, fmt.Errorf("invalid op: %s", program[i].op)
		}
	}
	return value, nil
}

func execute(program []instruction) (int, error) {
	value := 0
	seen := map[int]bool{}
	for i := 0; !seen[i]; {
		if i < 0 {
			return value, fmt.Errorf("invalid instruction number: %d", i)
		}
		if i == len(program) {
			return value, nil
		}
		seen[i] = true
		switch program[i].op {
		case nop:
			i++
		case acc:
			value += program[i].arg
			i++
		case jmp:
			i += program[i].arg
		}
	}
	return value, fmt.Errorf("program aborted: loop detected")
}

type instruction struct {
	op  operation
	arg int
}

type operation string

const (
	nop = operation("nop")
	acc = operation("acc")
	jmp = operation("jmp")
)

func parse(line string) (instruction, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return instruction{}, fmt.Errorf("expected two part instruction, got %s", line)
	}
	switch operation(parts[0]) {
	case nop, acc, jmp:
		arg, err := strconv.Atoi(parts[1])
		if err != nil {
			return instruction{}, fmt.Errorf("error converting instruction argument to int: %w", err)
		}
		return instruction{op: operation(parts[0]), arg: arg}, nil
	default:
		return instruction{}, fmt.Errorf("invalid op: %s", parts[0])
	}
}
