package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/danielmmetz/advent-of-code-2020/day01"
	"github.com/danielmmetz/advent-of-code-2020/day02"
	"github.com/danielmmetz/advent-of-code-2020/day03"
	"github.com/danielmmetz/advent-of-code-2020/day04"
	"github.com/danielmmetz/advent-of-code-2020/day05"
	"github.com/danielmmetz/advent-of-code-2020/day06"
	"github.com/danielmmetz/advent-of-code-2020/day07"
	"github.com/danielmmetz/advent-of-code-2020/day08"
	"github.com/danielmmetz/advent-of-code-2020/day09"
	"github.com/danielmmetz/advent-of-code-2020/day10"
	"github.com/danielmmetz/advent-of-code-2020/day11"
	"github.com/danielmmetz/advent-of-code-2020/day12"
	"github.com/danielmmetz/advent-of-code-2020/day13"
)

func main() {
	if err := mainE(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func mainE() error {
	var day int
	var part int
	flag.IntVar(&day, "day", 0, "the day's code to target for execution")
	flag.IntVar(&part, "part", 1, "1 or 2")
	flag.Parse()

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading stdin: %w", err)
	}

	var result string
	var err error

	p, err := aoc.ToPart(part)
	if err != nil {
		return err
	}
	switch day {
	case 1:
		result, err = day01.MainE(lines, p)
	case 2:
		result, err = day02.MainE(lines, p)
	case 3:
		result, err = day03.MainE(lines, p)
	case 4:
		result, err = day04.MainE(lines, p)
	case 5:
		result, err = day05.MainE(lines, p)
	case 6:
		result, err = day06.MainE(lines, p)
	case 7:
		result, err = day07.MainE(lines, p)
	case 8:
		result, err = day08.MainE(lines, p)
	case 9:
		result, err = day09.MainE(lines, p)
	case 10:
		result, err = day10.MainE(lines, p)
	case 11:
		result, err = day11.MainE(lines, p)
	case 12:
		result, err = day12.MainE(lines, p)
	case 13:
		result, err = day13.MainE(lines, p)
	default:
		err = fmt.Errorf("invalid day provided")
	}
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
