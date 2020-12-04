package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
)

func MainE(lines []string, part aoc.Part) (string, error) {
	passports, err := parse(lines)
	if err != nil {
		return "", fmt.Errorf("error parsing passports: %w", err)
	}
	var valid int
	for _, p := range passports {
		if p.Valid(part) {
			valid++
		}
	}
	return strconv.Itoa(valid), nil
}

func parse(lines []string) ([]passport, error) {
	if len(lines) == 0 {
		return []passport{}, nil
	}

	var passports []passport
	var p passport
	for _, line := range lines {
		if line == "" {
			passports = append(passports, p)
			p = passport{}
			continue
		}
		for _, component := range strings.Split(line, " ") {
			component := strings.TrimSpace(component)
			sepIdx := strings.Index(component, ":")
			if sepIdx == -1 {
				return passports, fmt.Errorf("failed to find separator for passport component: %s", component)
			}
			k, v := component[:sepIdx], component[sepIdx+1:]
			if err := p.Set(k, v); err != nil {
				return passports, err
			}
		}
	}
	passports = append(passports, p)
	return passports, nil
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) Set(key, value string) error {
	switch key {
	case "byr":
		p.byr = value
	case "iyr":
		p.iyr = value
	case "eyr":
		p.eyr = value
	case "hgt":
		p.hgt = value
	case "hcl":
		p.hcl = value
	case "ecl":
		p.ecl = value
	case "pid":
		p.pid = value
	case "cid":
		p.cid = value
	default:
		return fmt.Errorf("invalid key: %s", key)
	}
	return nil
}

func (p *passport) Valid(part aoc.Part) bool {
	switch part {
	case aoc.Part1:
		return p.p1Valid()
	case aoc.Part2:
		return p.p2Valid()
	default:
		return false
	}
}

func (p *passport) p1Valid() bool {
	required := []string{
		p.byr,
		p.iyr,
		p.eyr,
		p.hgt,
		p.hcl,
		p.ecl,
		p.pid,
	}
	for _, r := range required {
		if r == "" {
			return false
		}
	}
	return true
}

func (p *passport) p2Valid() bool {
	checks := []bool{
		withinInclusiveRange(p.byr, 1920, 2002),
		withinInclusiveRange(p.iyr, 2010, 2020),
		withinInclusiveRange(p.eyr, 2020, 2030),
		func() bool {
			if strings.HasSuffix(p.hgt, "cm") && withinInclusiveRange(p.hgt[:len(p.hgt)-2], 150, 193) {
				return true
			}
			if strings.HasSuffix(p.hgt, "in") && withinInclusiveRange(p.hgt[:len(p.hgt)-2], 59, 76) {
				return true
			}
			return false
		}(),
		isHexColor(p.hcl),
		isValidEyeColor(p.ecl),
		isValidPassportID(p.pid),
	}
	for _, check := range checks {
		if !check {
			return false
		}
	}
	return true
}

func withinInclusiveRange(candidate string, lower, upper int) bool {
	cast, err := strconv.Atoi(candidate)
	if err != nil {
		return false
	}
	return lower <= cast && cast <= upper
}

func isHexColor(s string) bool {
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}
	for _, c := range s[1:] {
		if !strings.ContainsRune("abcdef0123456789", c) {
			return false
		}
	}
	return true
}

func isValidEyeColor(c string) bool {
	switch c {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func isValidPassportID(id string) bool {
	if len(id) != 9 {
		return false
	}
	_, err := strconv.Atoi(id)
	return err == nil
}
