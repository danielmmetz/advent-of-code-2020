package day10

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/require"
)

func input1() []string {
	return []string{
		"16",
		"10",
		"15",
		"5",
		"1",
		"11",
		"7",
		"19",
		"6",
		"12",
		"4",
	}
}

func input2() []string {
	return []string{
		"28",
		"33",
		"18",
		"42",
		"31",
		"14",
		"46",
		"20",
		"48",
		"47",
		"24",
		"23",
		"49",
		"45",
		"19",
		"38",
		"39",
		"11",
		"1",
		"32",
		"25",
		"35",
		"8",
		"17",
		"7",
		"9",
		"4",
		"2",
		"34",
		"10",
		"3",
	}
}

func TestPart1(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{input1(), "35"},
		{input2(), "220"},
	}
	for _, c := range cases {
		t.Run(c.expected, func(t *testing.T) {
			result, err := MainE(c.input, aoc.Part1)
			require.NoError(t, err)
			require.Equal(t, c.expected, result)

		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{input1(), "8"},
		{input2(), "19208"},
	}
	for _, c := range cases {
		t.Run(c.expected, func(t *testing.T) {
			result, err := MainE(c.input, aoc.Part2)
			require.NoError(t, err)
			require.Equal(t, c.expected, result)

		})
	}
}
