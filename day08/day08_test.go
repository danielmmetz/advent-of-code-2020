package day08

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
}

func TestPart1(t *testing.T) {
	result, err := MainE(input(), aoc.Part1)
	require.NoError(t, err)
	require.Equal(t, "5", result)
}

func TestPart2(t *testing.T) {
	result, err := MainE(input(), aoc.Part2)
	require.NoError(t, err)
	require.Equal(t, "8", result)
}
