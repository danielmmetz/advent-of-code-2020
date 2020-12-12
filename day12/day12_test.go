package day12

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
}

func TestPart1(t *testing.T) {
	result, err := MainE(input(), aoc.Part1)
	require.NoError(t, err)
	require.Equal(t, "25", result)
}

func TestPart2(t *testing.T) {
	result, err := MainE(input(), aoc.Part2)
	require.NoError(t, err)
	require.Equal(t, "286", result)
}
