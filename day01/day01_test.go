package day01

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}
}

func TestPart1(t *testing.T) {
	output, err := MainE(input(), aoc.Part1)
	require.NoError(t, err)
	assert.Equal(t, "514579", output)
}

func TestPart2(t *testing.T) {
	output, err := MainE(input(), aoc.Part2)
	require.NoError(t, err)
	assert.Equal(t, "241861950", output)
}
