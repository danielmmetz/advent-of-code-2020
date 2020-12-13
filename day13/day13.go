package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	switch part {
	case aoc.Part1:
		earliestTime, buses, err := parse1(lines)
		if err != nil {
			return "", fmt.Errorf("error parsing input: %w", err)
		}
		return strconv.Itoa(part1(earliestTime, buses)), nil
	case aoc.Part2:
		if len(lines) != 2 {
			return "", fmt.Errorf("unexpected input length: %d", len(lines))
		}
		buses, err := parse2(lines[1])
		if err != nil {
			return "", fmt.Errorf("error parsing input: %w", err)
		}
		return strconv.Itoa(firstSequentialTime(buses)), nil
	default:
		return "", fmt.Errorf("invalid part specified: %v", part)
	}
}

func parse1(lines []string) (int, []int, error) {
	if len(lines) != 2 {
		return 0, nil, fmt.Errorf("unexpected input length: %d", len(lines))
	}
	minTime, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, fmt.Errorf("invalid min time: %w", err)
	}
	var busIntervals []int
	buses := strings.Split(lines[1], ",")
	for i, bus := range buses {
		if bus == "x" {
			continue
		}
		t, err := strconv.Atoi(bus)
		if err != nil {
			return 0, nil, fmt.Errorf("invalid time for bus %d: %w", i, err)
		}
		busIntervals = append(busIntervals, t)
	}
	return minTime, busIntervals, nil
}

func part1(minTime int, buses []int) int {
	earliestTime := minTime * 1000
	earliestBus := 0
	for _, bus := range buses {
		min := ceilDiv(minTime, bus) * bus
		if min < earliestTime {
			earliestTime = min
			earliestBus = bus
		}
	}
	return earliestBus * (earliestTime - minTime)
}

func divMod(numerator, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

func ceilDiv(numerator, denominator int) int {
	div, rem := divMod(numerator, denominator)
	if rem == 0 {
		return div
	}
	return div + 1
}

type bus struct {
	id   int
	wild bool
}

func parse2(line string) ([]bus, error) {
	busCandidates := strings.Split(line, ",")

	var buses []bus
	for i, b := range busCandidates {
		if b == "x" {
			buses = append(buses, bus{wild: true})
			continue
		}
		t, err := strconv.Atoi(b)
		if err != nil {
			return buses, fmt.Errorf("invalid time for bus %d: %w", i, err)
		}
		buses = append(buses, bus{id: t})
	}
	return buses, nil
}

func firstSequentialTime(buses []bus) int {
	switch {
	case len(buses) == 0:
		return 0
	case len(buses) == 1:
		if buses[0].wild {
			return 0
		}
		return buses[0].id
	}

	time := 0
	for {
		steps := map[int]struct{}{}
		for i, bus := range buses {
			if bus.wild {
				continue
			}
			if (time+i)%bus.id != 0 {
				break
			}
			steps[bus.id] = struct{}{}
		}
		if isSequential(buses, time) {
			return time
		}
		step := 1
		for k := range steps {
			step *= k
		}
		time += step
	}
}

func isSequential(buses []bus, time int) bool {
	for i, bus := range buses {
		if bus.wild {
			continue
		}
		if (time+i)%bus.id != 0 {
			return false
		}
	}
	return true
}
