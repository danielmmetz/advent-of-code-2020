package day14

import (
	"fmt"
	"testing"

	"github.com/danielmmetz/advent-of-code-2020/aoc"
	"github.com/stretchr/testify/require"
)

func input1() []string {
	return []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
}

func input2() []string {
	return []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
}

func TestPart1(t *testing.T) {
	result, err := MainE(input1(), aoc.Part1)
	require.NoError(t, err)
	require.Equal(t, "165", result)
}

func TestPart2(t *testing.T) {
	result, err := MainE(input2(), aoc.Part2)
	require.NoError(t, err)
	require.Equal(t, "208", result)
}

func TestBitmapRender1(t *testing.T) {
	cases := []struct {
		input, expected int
	}{
		{11, 73},
		{101, 101},
		{0, 64},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.input), func(t *testing.T) {
			result := bitmask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X").render1(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}

func TestBitmapRender2(t *testing.T) {
	cases := []struct {
		mask     string
		location int
		results  []int
	}{
		{"000000000000000000000000000000X1001X", 42, []int{26, 27, 58, 59}},
		{"00000000000000000000000000000000X0XX", 26, []int{16, 17, 18, 19, 24, 25, 26, 27}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.location), func(t *testing.T) {
			results := bitmask(c.mask).render2(c.location)
			require.Equal(t, c.results, results)
		})
	}
}
