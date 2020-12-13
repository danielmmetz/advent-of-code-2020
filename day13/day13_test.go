package day13

import (
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/require"
)

func input() []string {
	return []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
}

func TestPart1(t *testing.T) {
	result, err := MainE(input(), aoc.Part1)
	require.NoError(t, err)
	require.Equal(t, "295", result)
}

func TestPart2(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"17,x,13,19", "3417"},
		{"67,7,59,61", "754018"},
		{"67,x,7,59,61", "779210"},
		{"67,7,x,59,61", "1261476"},
		{"1789,37,47,1889", "1202161486"},
	}
	for _, c := range cases {
		t.Run(c.expected, func(t *testing.T) {
			result, err := MainE([]string{"", c.input}, aoc.Part2)
			require.NoError(t, err)
			require.Equal(t, c.expected, result)

		})
	}
}
