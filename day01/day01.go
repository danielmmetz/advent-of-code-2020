package day01

import (
	"fmt"
	"strconv"
)

func MainE(input []string, part int) (string, error) {
	var fn func([]int) ([]int, error)
	switch part {
	case 1:
		fn = findPair
	case 2:
		fn = findTriplet
	default:
		return "", fmt.Errorf("invalid part number specified: %d", part)
	}
	var ints []int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			return "", fmt.Errorf("error parsing input: expected int, got %s: %w", s, err)
		}
		ints = append(ints, i)
	}
	entries, err := fn(ints)
	if err != nil {
		return "", err
	}
	return format(entries...), nil
}

func findPair(candidates []int) ([]int, error) {
	for i, a := range candidates[:len(candidates)-1] {
		for _, b := range candidates[i+1:] {
			if a+b == 2020 {
				return []int{a, b}, nil
			}
		}
	}
	return nil, fmt.Errorf("no valid pair amongst candidates")
}

func findTriplet(candidates []int) ([]int, error) {
	for i, a := range candidates {
		for j, b := range candidates[i+1:] {
			if i == j {
				continue
			}
			for k, c := range candidates[j+1:] {
				if j == k {
					continue
				}
				if a+b+c == 2020 {
					return []int{a, b, c}, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("no valid pair amongst candidates")
}

func format(ints ...int) string {
	if len(ints) == 0 {
		return "0"
	}
	result := 1
	for _, i := range ints {
		result *= i
	}
	return strconv.Itoa(result)
}
