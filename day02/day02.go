package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

var rulePattern = regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]*)`)

func MainE(lines []string, part aoc.Part) (string, error) {
	var rules []rule
	var passwords []string
	for i, line := range lines {
		rule, candidate, err := parse(line)
		if err != nil {
			return "", fmt.Errorf("error encountered processing line %d: %w", i+1, err)
		}
		rules = append(rules, rule)
		passwords = append(passwords, candidate)
	}

	var valid int
	switch part {
	case aoc.Part1:
		for i := range rules {
			if rules[i].p1validates(passwords[i]) {
				valid++
			}
		}
		return strconv.Itoa(valid), nil
	case aoc.Part2:
		for i := range rules {
			if rules[i].p2validates(passwords[i]) {
				valid++
			}
		}
		return strconv.Itoa(valid), nil
	default:
		return "", fmt.Errorf("invalid part number specified: %d", part)
	}
}

type rule struct {
	min, max  int
	character string
}

func (r rule) p1validates(password string) bool {
	count := strings.Count(password, r.character)
	return r.min <= count && count <= r.max
}

func (r rule) p2validates(password string) bool {
	if len(password) < r.max {
		return false
	}
	cond1 := string(password[r.min-1]) == r.character
	cond2 := string(password[r.max-1]) == r.character
	return cond1 != cond2
}

func parse(input string) (rule, string, error) {
	submatch := rulePattern.FindStringSubmatch(input)
	if len(submatch) == 0 {
		return rule{}, "", fmt.Errorf("input did not match expected format: %s", input)
	}
	min, _ := strconv.Atoi(submatch[1]) // valid by definition of match
	max, _ := strconv.Atoi(submatch[2]) // valid by definition of match
	return rule{min: min, max: max, character: submatch[3]}, submatch[4], nil
}
