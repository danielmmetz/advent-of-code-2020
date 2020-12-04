package day03

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
}

func TestPart1(t *testing.T) {
	result, err := MainE(input(), aoc.Part1)
	require.NoError(t, err)
	assert.Equal(t, "7", result)
}

func TestPart2(t *testing.T) {
	result, err := MainE(input(), aoc.Part2)
	require.NoError(t, err)
	assert.Equal(t, "336", result)
}
