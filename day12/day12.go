package day12

import (
	"fmt"
	"strconv"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	switch part {
	case aoc.Part1:
		s := ship1{facing: east}
		if err := process(&s, lines); err != nil {
			return "", fmt.Errorf("error processing instructions: %w", err)
		}
		return strconv.Itoa(manhattanDistanceFromOrigin(s.x, s.y)), nil
	case aoc.Part2:
		s := ship2{w: waypoint{x: 10, y: 1}}
		if err := process(&s, lines); err != nil {
			return "", fmt.Errorf("error processing instructions: %w", err)
		}
		return strconv.Itoa(manhattanDistanceFromOrigin(s.x, s.y)), nil
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

func process(p processor, instructions []string) error {
	return evaluator{p: p}.process(instructions)
}

type evaluator struct {
	p processor
}

func (e evaluator) process(instructions []string) error {
	for i, line := range instructions {
		if err := e.p.process(line); err != nil {
			return fmt.Errorf("error following instruction %d: %w", i, err)
		}
	}
	return nil
}

type processor interface {
	process(step string) error
}

type ship1 struct {
	x, y   int
	facing direction
}

func (s *ship1) process(step string) error {
	if len(step) < 2 {
		return fmt.Errorf("bad instruction: %s", step)
	}
	magnitude, err := strconv.Atoi(step[1:])
	if err != nil {
		return fmt.Errorf("unable to parse magnitude from %s: %w", step, err)
	}
	switch step[0] {
	case 'N':
		s.y += magnitude
	case 'S':
		s.y -= magnitude
	case 'E':
		s.x += magnitude
	case 'W':
		s.x -= magnitude
	case 'R':
		return s.rotate(magnitude)
	case 'L':
		return s.rotate(360 - magnitude)
	case 'F':
		return s.process(fmt.Sprintf("%v%d", s.facing, magnitude))
	}
	return nil
}

func (s *ship1) rotate(magnitude int) error {
	if magnitude%90 != 0 {
		return fmt.Errorf("invalid rotational magntidue specified: %d", magnitude)
	}
	idx, err := index(directions, s.facing)
	if err != nil {
		return fmt.Errorf("facing invalid direction: %v", s.facing)
	}
	increments := magnitude / 90 % 4
	s.facing = directions[(idx+increments)%4]
	return nil
}

type ship2 struct {
	x, y int
	w    waypoint
}

type waypoint struct {
	x, y int
}

func (s *ship2) process(step string) error {
	if len(step) < 2 {
		return fmt.Errorf("bad instruction: %s", step)
	}
	magnitude, err := strconv.Atoi(step[1:])
	if err != nil {
		return fmt.Errorf("unable to parse magnitude from %s: %w", step, err)
	}
	switch step[0] {
	case 'N':
		s.w.y += magnitude
	case 'S':
		s.w.y -= magnitude
	case 'E':
		s.w.x += magnitude
	case 'W':
		s.w.x -= magnitude
	case 'R':
		return s.rotateWaypoint(magnitude)
	case 'L':
		return s.rotateWaypoint(360 - magnitude)
	case 'F':
		s.x += s.w.x * magnitude
		s.y += s.w.y * magnitude
	}
	return nil
}

func (s *ship2) rotateWaypoint(magnitude int) error {
	if magnitude%90 != 0 {
		return fmt.Errorf("invalid rotational magntidue specified: %d", magnitude)
	}
	increments := magnitude / 90 % 4
	for i := 0; i < increments; i++ {
		s.w.x, s.w.y = s.w.y, -s.w.x
	}
	return nil
}

type direction rune

func (d direction) String() string {
	return string(d)
}

const (
	north direction = 'N'
	south direction = 'S'
	east  direction = 'E'
	west  direction = 'W'
)

var directions = []direction{north, east, south, west}

func index(haystack []direction, needle direction) (int, error) {
	for i, candidate := range haystack {
		if candidate == needle {
			return i, nil
		}
	}
	return 0, fmt.Errorf("not found")
}

func manhattanDistanceFromOrigin(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}
