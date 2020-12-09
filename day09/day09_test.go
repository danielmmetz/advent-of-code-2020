package day09

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"35",
		"20",
		"15",
		"25",
		"47",
		"40",
		"62",
		"55",
		"65",
		"95",
		"102",
		"117",
		"150",
		"182",
		"127",
		"219",
		"299",
		"277",
		"309",
		"576",
	}
}

func TestPart1(t *testing.T) {
	result, err := MainE(input(), aoc.Part1, preambleLength(5))
	require.NoError(t, err)
	require.Equal(t, "127", result)
}

func TestPart2(t *testing.T) {
	result, err := MainE(input(), aoc.Part2, preambleLength(5))
	require.NoError(t, err)
	require.Equal(t, "62", result)
}
