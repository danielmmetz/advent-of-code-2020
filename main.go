package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/danielmmetz/advent-of-code-2020/day01"
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

	switch day {
	case 1:
		result, err = day01.MainE(lines, part)
	default:
		err = fmt.Errorf("invalid day provided")
	}
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
