package day11

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
}

func TestPart1(t *testing.T) {
	result, err := MainE(input(), aoc.Part1)
	require.NoError(t, err)
	require.Equal(t, "37", result)
}

func TestPart2(t *testing.T) {
	result, err := MainE(input(), aoc.Part2)
	require.NoError(t, err)
	require.Equal(t, "26", result)
}
