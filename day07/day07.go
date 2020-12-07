package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

var rulePatternContains = regexp.MustCompile(`(?P<key>\w+\s\w+) bags contain(?:(?:,? )(?P<count>\d+) (?P<value>\w+\s\w+) bags?)+.*\.`)
var rulePatternContainsPrefix = regexp.MustCompile(`(?P<key>\w+\s\w+) bags contain`)
var rulePatternContainsSuffix = regexp.MustCompile(`(?P<count>\d+) (?P<value>\w+\s\w+) bags?`)
var rulePatternEmpty = regexp.MustCompile(`(?P<key>\w+\s\w+) bags contain no other bags\.`)

func MainE(lines []string, part aoc.Part) (string, error) {
	var rules []rule
	for _, line := range lines {
		r, err := parseLine(line)
		if err != nil {
			return "", fmt.Errorf("error parsing rule: %w", err)
		}
		rules = append(rules, r)
	}
	rs := ruleSet{Rules: rules}
	switch part {
	case aoc.Part1:
		count, err := rs.howManyCanContain("shiny gold")
		if err != nil {
			return "", fmt.Errorf("error calculating how many bags may contain target: %w", err)
		}
		return strconv.Itoa(count), nil
	case aoc.Part2:
		count, err := rs.countContained("shiny gold")
		if err != nil {
			return "", fmt.Errorf("error calculating how many bags target contains: %w", err)
		}
		return strconv.Itoa(count), nil
	default:
		return "", fmt.Errorf("invalid part: %v", part)
	}
}

type ruleSet struct {
	Rules []rule

	ruleMap        map[string]rule
	ruleResolution map[string]map[string]int
}

func (rs *ruleSet) init() error {
	if rs.ruleMap == nil {
		rs.ruleMap = map[string]rule{}
	}
	if rs.ruleResolution == nil {
		rs.ruleResolution = map[string]map[string]int{}
	}

	for _, r := range rs.Rules {
		rs.ruleMap[r.color] = r
	}
	for _, r := range rs.Rules {
		if err := rs.resolveRule(r); err != nil {
			return fmt.Errorf("error resolving rule for %s: %w", r.color, err)
		}
	}
	return nil
}

func (rs *ruleSet) howManyCanContain(target string) (int, error) {
	if err := rs.init(); err != nil {
		return 0, err
	}

	var hits int
	for _, contains := range rs.ruleResolution {
		if contains[target] > 0 {
			hits++
		}
	}
	return hits, nil
}

func (rs *ruleSet) countContained(target string) (int, error) {
	if err := rs.init(); err != nil {
		return 0, err
	}

	targetRule := rs.ruleMap[target]

	hits := 0
	for _, contained := range targetRule.values {
		containedCount, err := rs.countContained(contained.color)
		if err != nil {
			return 0, fmt.Errorf("error determining count of component bag %s: %w", contained.color, err)
		}
		hits += contained.count * (1 + containedCount)
	}
	return hits, nil
}

func (rs *ruleSet) resolveRule(r rule) error {
	if _, ok := rs.ruleResolution[r.color]; ok {
		return nil
	}
	rs.ruleResolution[r.color] = map[string]int{}
	for _, containedValue := range r.values {
		rs.ruleResolution[r.color][containedValue.color] = containedValue.count
		if err := rs.resolveRule(rs.ruleMap[containedValue.color]); err != nil {
			return fmt.Errorf("error resolving rule for %s: %w", containedValue.color, err)
		}
		for k := range rs.ruleResolution[containedValue.color] {
			rs.ruleResolution[r.color][k] = containedValue.count
		}
	}
	return nil
}

type rule struct {
	color  string
	values []value
}

type value struct {
	count int
	color string
}

func parseLine(line string) (rule, error) {
	if match := rulePatternEmpty.FindStringSubmatch(line); match != nil {
		return rule{color: match[1]}, nil
	}
	if fullMatch := rulePatternContains.FindAllStringSubmatch(line, -1); fullMatch != nil {
		prefix := rulePatternContainsPrefix.FindStringSubmatch(line)
		r := rule{color: prefix[1]}
		suffix := strings.Split(line, "contain")[1]
		for _, match := range rulePatternContainsSuffix.FindAllStringSubmatch(suffix, -1) {
			v := value{color: match[2]}
			v.count, _ = strconv.Atoi(match[1])
			r.values = append(r.values, v)
		}
		return r, nil
	}
	return rule{}, fmt.Errorf("no pattern matched")
}
