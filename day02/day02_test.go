package day02

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
}

func TestPart1(t *testing.T) {
	count, err := MainE(input(), aoc.Part1)
	require.NoError(t, err)
	assert.Equal(t, "2", count)
}

func TestPart2(t *testing.T) {
	count, err := MainE(input(), aoc.Part2)
	require.NoError(t, err)
	assert.Equal(t, "1", count)
}
